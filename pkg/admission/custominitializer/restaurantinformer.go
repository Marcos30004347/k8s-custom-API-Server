package custominitializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/informers/externalversions"
)

type bazInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = bazInformerPluginInitializer{}

// New creates an instance of custom admission plugins initializer.
func New(informers informers.SharedInformerFactory) bazInformerPluginInitializer {
	return bazInformerPluginInitializer{
		informers: informers,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data
func (i bazInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsbazInformerFactory); ok {
		wants.SetBazInformerFactory(i.informers)
	}
}
