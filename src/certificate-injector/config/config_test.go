package config_test

import (
	"certificate-injector/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container", func() {
	var (
		conf     config.Container
		certData []byte
	)
	BeforeEach(func() {
		conf = config.NewContainer()
		certData = []byte("really-a-cert")
	})

	It("encodes a script to import the certificates and writes it to config.json", func() {
		err := conf.Write(certData)
		Expect(err).NotTo(HaveOccurred())

		data, err := ioutil.ReadFile("config.json")
		Expect(err).NotTo(HaveOccurred())

		container := config.ContainerJSON{}
		json.Unmarshal(data, &container)
		Expect(container.Process.Cwd).To(Equal("C:\\"))

		decoded, err := base64.StdEncoding.DecodeString(container.Process.Args[2])
		Expect(err).NotTo(HaveOccurred())
		Expect(string(decoded)).To(Equal(fmt.Sprintf(config.ImportCertificatePs, string(certData))))
	})
})
