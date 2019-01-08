package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
}

type config struct {
	Process process `json:"process"`
}

type process struct {
	Args []string `json:"args"`
	Cwd  string   `json:"cwd"`
}

func NewConfig() Config {
	return Config{}
}

// Creates a powershell script to write the certs
// to a file and import the certificate. It appends
// this script as a process to a config.json that will
// be run on the container.
func (c Config) Write() error {

	// encodedCertData := base64.StdEncoding.EncodeToString(certData)
	// addCertScript := fmt.Sprintf(`
	// $ErrorActionPreference = "Stop";
	// trap { $host.SetShouldExit(1) }
	// $certFile=[System.IO.Path]::GetTempFileName()
	// $decodedCertData = [Convert]::FromBase64String("%s")
	// [IO.File]::WriteAllBytes($certFile, $decodedCertData)
	// Import-Certificate -CertStoreLocation Cert:\\LocalMachine\Root -FilePath $certFile
	// Remove-Item $certFile
	// `, encodedCertData)
	// encodedScript := base64.StdEncoding.EncodeToString([]byte(addCertScript))

	conf := config{
		Process: process{
			Args: []string{"powershell.exe", "-EncodedCommand", ""},
			Cwd:  `C:\`,
		},
	}
	configJson, _ := json.Marshal(conf)

	err := ioutil.WriteFile("config.json", configJson, 0644)
	if err != nil {
		return fmt.Errorf("Write config.json failed: %s", err)
	}

	return nil
}
