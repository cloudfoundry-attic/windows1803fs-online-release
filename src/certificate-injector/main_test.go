package main_test

import (
	. "certificate-injector"
	fakes "certificate-injector/fakes"
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("certificate-injector", func() {
	var (
		fakeUtil *fakes.Util
		fakeCmd  *fakes.Cmd
		args     []string
	)

	BeforeEach(func() {
		fakeUtil = &fakes.Util{}
		fakeCmd = &fakes.Cmd{}
	})

	Context("when the layer has the layerAdded annotation", func() {
		BeforeEach(func() {
			args = []string{"certificate-injector.exe", "", "", "first-image-uri"}
			fakeUtil.ContainsHydratorAnnotationCall.Returns.Contains = true
		})

		It("calls hydrator to remove the custom layer", func() {
			_ = Run(args, fakeUtil, fakeCmd)
			Expect(fakeCmd.RunCall.CallCount).To(Equal(1))
			Expect(fakeCmd.RunCall.Receives.Executable).To(ContainSubstring("hydrate.exe"))
			Expect(fakeCmd.RunCall.Receives.Args).To(ContainElement("first-image-uri"))
		})

		Context("when hydrator fails to remove the custom layer", func() {
			BeforeEach(func() {
				fakeCmd.RunCall.Returns.Error = errors.New("hydrator is unhappy")
			})
			It("should return a helpful error", func() {
				err := Run(args, fakeUtil, fakeCmd)
				Expect(err).To(MatchError("hydrate.exe remove-layer failed: hydrator is unhappy\n"))
			})
		})
	})

	Context("when the layer does not have the layerAdded annotation", func() {
		BeforeEach(func() {
			args = []string{"certificate-injector.exe", "", "", "first-image-uri"}
		})

		It("does not call hydrator to remove the custom layer", func() {
			err := Run(args, fakeUtil, fakeCmd)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeUtil.ContainsHydratorAnnotationCall.CallCount).To(Equal(1))
			Expect(fakeUtil.ContainsHydratorAnnotationCall.Receives.OCIImagePath).To(Equal("first-image-uri"))
			Expect(fakeCmd.RunCall.CallCount).To(Equal(0))
		})
	})

	Describe("cert_file", func() {
		Context("when the cert_file does not exist", func() {
			BeforeEach(func() {
				args = []string{"certificate-injector.exe", "", "not-a-real-file.crt", "first-image-uri"}
			})

			It("returns a helpful error", func() {
				err := Run(args, fakeUtil, fakeCmd)
				Expect(err).To(MatchError("Failed to read cert_file: open not-a-real-file.crt: no such file or directory"))
			})
		})

		Context("when there are no trusted certs to inject", func() {
			BeforeEach(func() {
				args = []string{"certificate-injector.exe", "", "fakes/empty.crt", "first-image-uri"}
			})

			It("does not check other arguments and exits successfully", func() {
				err := Run(args, fakeUtil, fakeCmd)
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeUtil.ContainsHydratorAnnotationCall.CallCount).To(Equal(0))
			})
		})
	})

	Context("when called with incorrect arguments", func() {
		It("returns a helpful error message with usage", func() {
			args := []string{"certificate-injector.exe"}
			err := Run(args, fakeUtil, fakeCmd)
			Expect(err).To(MatchError(fmt.Sprintf("usage: %s <driver_store> <cert_file> <image_uri>...\n", args[0])))
		})
	})
})
