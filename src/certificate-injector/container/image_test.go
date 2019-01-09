package container_test

import (
	"certificate-injector/container"
	"certificate-injector/container/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image", func() {
	It("checks if a layer contains a hydrator annotation", func() {
		handler := &fakes.Handler{}
		image := container.NewImage(handler)

		Expect(image.ContainsHydratorAnnotation("imageUri")).To(BeTrue())
		Expect(handler.ReadMetadataCall.CallCount).To(Equal(1))
	})
})
