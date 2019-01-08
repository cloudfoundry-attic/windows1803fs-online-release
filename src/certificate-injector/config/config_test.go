package config_test

import (
	"certificate-injector/config"
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var conf config.Config
	BeforeEach(func() {
		conf = config.NewConfig()
	})

	It("encodes a script to import the certificates and writes it to config.json", func() {
		err := conf.Write()
		Expect(err).NotTo(HaveOccurred())

		data, err := ioutil.ReadFile("config.json")
		Expect(err).NotTo(HaveOccurred())

		jsonConfig := config.ConfigJson{}
		json.Unmarshal(data, &jsonConfig)
		Expect(jsonConfig.Process.Cwd).To(Equal("C:\\"))
		Expect(jsonConfig.Process.Args[2]).To(Equal(""))
	})
})
