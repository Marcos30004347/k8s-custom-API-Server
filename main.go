package main

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/Marcos30004347/k8s-custom-API-Server/pkg/cmd/server"
	"github.com/Marcos30004347/k8s-custom-API-Server/pkg/controller"

	"k8s.io/component-base/logs"
	"k8s.io/klog"

	clientset "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/clientset/versioned"
	informers "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/informers/externalversions"
	kubeinformers "k8s.io/client-go/informers"
)

var (
	stopCh        <-chan struct{}
	apiController *controller.Controller
	command       *cobra.Command
)

func runController(wg *sync.WaitGroup) {
	defer wg.Done()

	if err := apiController.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func runServerAPI(wg *sync.WaitGroup) {
	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}

func main() {
	// Read command line flags
	kubeconfig := flag.String("kubeconfig", "", "a string")
	master := flag.String("master", "", "a string")
	certdir := flag.String("cert-dir", "", "a string")
	etcd := flag.String("etcd-servers", "", "a string")
	authenticationc := flag.String("authentication-kubeconfig", "", "a string")
	authorizationk := flag.String("authorization-kubeconfig", "", "a string")
	secureport := flag.String("secure-port", "", "a string")

	klog.Info("Using secure port in: %s", *secureport)
	klog.Info("Using certificate directory in: %s", *certdir)
	klog.Info("Using kubeconfig in: %s", *kubeconfig)
	klog.Info("Using etcd cluster in: %s", *etcd)
	klog.Info("Using authentication-kubeconfig in: %s", *authenticationc)
	klog.Info("Using authorization-kubeconfig in: %s", *authorizationk)

	flag.Parse()

	// Init logs
	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh = genericapiserver.SetupSignalHandler()
	options := server.NewCustomServerOptions(os.Stdout, os.Stderr)

	command = server.NewCommandStartCustomServer(options, stopCh)
	command.Flags().AddGoFlagSet(flag.CommandLine)

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)

	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)

	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	sharedInformer := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	apiController = controller.NewController(kubeClient, exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		sharedInformer.Baz().V1alpha1().Foos())

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)
	sharedInformer.Start(stopCh)

	var wg sync.WaitGroup

	wg.Add(2)

	go runServerAPI(&wg)
	go runController(&wg)

	wg.Wait()
}
