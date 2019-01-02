package fakes

import "fmt"

type Util struct {
	containsHydratorAnnotation bool
}

func (u *Util) ContainsHydratorAnnotationReturns(annotations bool) {
	u.containsHydratorAnnotation = annotations
}

func (u *Util) ContainsHydratorAnnotation(ociImagePath string) bool {
	fmt.Printf("SSUP\n")
	return u.containsHydratorAnnotation
}
