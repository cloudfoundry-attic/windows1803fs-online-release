package fakes

type Util struct {
	ContainsHydratorAnnotationCall struct {
		CallCount int
		Receives  struct {
			OCIImagePath string
		}
		Returns struct {
			Contains bool
		}
	}
}

func (u *Util) ContainsHydratorAnnotation(ociImagePath string) bool {
	u.ContainsHydratorAnnotationCall.CallCount++
	u.ContainsHydratorAnnotationCall.Receives.OCIImagePath = ociImagePath

	return u.ContainsHydratorAnnotationCall.Returns.Contains
}
