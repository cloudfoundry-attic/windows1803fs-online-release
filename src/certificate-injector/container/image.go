package container

import (
	oci "github.com/opencontainers/image-spec/specs-go/v1"
)

type Image struct {
	handler handler
}

type handler interface {
	ReadMetadata() (oci.Manifest, oci.Image, error)
}

func NewImage(handler handler) Image {
	return Image{
		handler: handler,
	}
}

func (u Image) ContainsHydratorAnnotation(path string) bool {
	// TODO: How does one find the annotation for an oci image?
	u.handler.ReadMetadata()
	return true
}
