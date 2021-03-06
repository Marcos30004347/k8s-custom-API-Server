package v1alpha1

import (
	"github.com/Marcos30004347/k8s-custom-API-Server/pkg/apis/baz"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_v1alpha1_FooSpec_To_baz_FooSpec is an autogenerated conversion function.
func Convert_v1alpha1_FooSpec_To_baz_FooSpec(in *FooSpec, out *baz.FooSpec, s conversion.Scope) error {
	idx := map[string]int{}

	for _, bar := range in.Bar {
		if i, duplicate := idx[bar]; duplicate {
			out.Bar[i].Quantity++
			continue
		}

		idx[bar] = len(out.Bar)

		out.Bar = append(out.Bar, baz.FooBar{
			Name:     bar,
			Quantity: 1,
		})
	}

	return nil
}

// Convert_baz_FooSpec_To_v1alpha1_FooSpec is an autogenerated conversion function.
func Convert_baz_FooSpec_To_v1alpha1_FooSpec(in *baz.FooSpec, out *FooSpec, s conversion.Scope) error {
	for i := range in.Bar {
		for j := 0; j < in.Bar[i].Quantity; j++ {
			out.Bar = append(out.Bar, in.Bar[i].Name)
		}
	}

	return nil
}
