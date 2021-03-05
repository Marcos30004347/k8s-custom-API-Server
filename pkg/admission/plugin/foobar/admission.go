package foobar

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"

	"github.com/Marcos30004347/k8s-custom-API-Server/pkg/admission/custominitializer"
	"github.com/Marcos30004347/k8s-custom-API-Server/pkg/apis/baz"

	informers "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/informers/externalversions"
	listers "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/listers/baz/v1alpha1"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("FooBar", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

// The Plugin structure
type Plugin struct {
	*admission.Handler
	barLister listers.BarLister
}

var _ = custominitializer.WantsbazInformerFactory(&Plugin{})

// Admit ensures that the object in-flight is of kind Foo.
// In addition checks that the bar are known.
func (d *Plugin) Admit(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != baz.Kind("Foo") {
		return nil
	}

	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	obj := a.GetObject()

	foo := obj.(*baz.Foo)

	for _, bar := range foo.Spec.Bar {
		if _, err := d.barLister.Get(a.GetNamespace() + "/" + bar.Name); err != nil && errors.IsNotFound(err) {
			return errors.NewForbidden(
				a.GetResource().GroupResource(),
				a.GetName(),
				fmt.Errorf("unknown bar: %s", bar.Name),
			)
		}
	}

	return nil
}

// SetBazInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Bar.
func (d *Plugin) SetBazInformerFactory(f informers.SharedInformerFactory) {
	d.barLister = f.Baz().V1alpha1().Bars().Lister()
	d.SetReadyFunc(f.Baz().V1alpha1().Bars().Informer().HasSynced)
}

// ValidateInitialization checks whether the plugin was correctly initialized.
func (d *Plugin) ValidateInitialization() error {
	if d.barLister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}

// New creates a new ban foo topping admission plugin
func New() (*Plugin, error) {
	return &Plugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
