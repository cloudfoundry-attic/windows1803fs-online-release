package main

type Util struct {
}

func NewUtil() Util {
	return Util{}
}

func (u Util) ContainsHydratorAnnotation(ociImage string) bool {
	return true
}
