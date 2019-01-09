package main

type Util struct {
}

func NewUtil() Util {
	return Util{}
}

func (u Util) ContainsHydratorAnnotation(ociImage string) bool {
	// TODO: How does one find the annotation for an oci image?
	return true
}
