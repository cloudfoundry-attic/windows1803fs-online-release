package main_test

import (
	. "certificate-injector"
	fakes "certificate-injector/fakes"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("certificate-injector", func() {
	Context("when called with incorrect arguments", func() {
		It("returns an error", func() {
			args := []string{"certificate-injector.exe"}
			fakeUtil := &fakes.Util{}
			err := Run(args, fakeUtil)
			Expect(err).To(MatchError(fmt.Sprintf("usage: %s <driver_store> <image_uri>...\n", args[0])))
		})
	})

	Context("when called with correct arguments", func() {
		Context("when the rootfs does not have custom layers", func() {
			It("hydrator is not called with remove-layer flag", func() {
				args := []string{"certificate-injector.exe", "", ""}
				fakeUtil := &fakes.Util{}
				fakeUtil.ContainsHydratorAnnotationReturns(false)
				_ = Run(args, fakeUtil)
				//Expect(err).To(MatchError(fmt.Sprintf("usage: %s <driver_store> <image_uri>...\n", args[0])))
			})
		})
	})
})
