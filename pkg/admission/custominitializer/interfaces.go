package custominitializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/Marcos30004347/k8s-custom-API-Server/pkg/generated/informers/externalversions"
)

// WantsbazInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsbazInformerFactory interface {
	SetBazInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
