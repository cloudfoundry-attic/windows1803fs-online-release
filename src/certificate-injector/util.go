package main

type Util struct {
}

func (u *Util) ContainsHydratorAnnotation(ociImage string) bool {
	return true
}
