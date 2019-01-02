package main

import (
	"fmt"
	"log"
	"os"
	/*"bytes"
	  "code.cloudfoundry.org/filelock"
	  "encoding/base64"
	  "encoding/json"
	  "io/ioutil"
	  "os/exec"
	  "path/filepath"
	  "strings"
	  "time"*/)

const (
	LockFileName    = "GrootRootfsMutex"
	grootBin        = "c:\\var\\vcap\\packages\\groot\\groot.exe"
	wincBin         = "c:\\var\\vcap\\packages\\winc\\winc.exe"
	diffExporterBin = "c:\\var\\vcap\\packages\\diff-exporter\\diff-exporter.exe"
	hydrateBin      = "c:\\var\\vcap\\packages\\hydrate\\hydrate.exe"
)

type UtilInterface interface {
	ContainsHydratorAnnotation(ociImagePath string) bool
}

func Run(args []string, util UtilInterface) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: %s <driver_store> <image_uri>...\n", args[0])
	}
	util.ContainsHydratorAnnotation(args[2])
	/*grootDriverStore := args[1]
	grootImageUris := args[2:]

	lock, err := filelock.NewLocker(filepath.Join(os.TempDir(), LockFileName)).Open()
	if err != nil {
		return fmt.Errorf("open lock: %s\n", err)
	}
	defer lock.Close()

	// inject some certificates
	certData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("%s\n", "Cannot read certificates")
	}


	if len(certData) == 0 {
		return nil
	}

	// creating an "add certificate" script that will be run inside a container
	encodedCertData := base64.StdEncoding.EncodeToString(certData)
	addCertScript := fmt.Sprintf(`
	$ErrorActionPreference = "Stop";
	trap { $host.SetShouldExit(1) }
	$certFile=[System.IO.Path]::GetTempFileName()
	$decodedCertData = [Convert]::FromBase64String("%s")
	[IO.File]::WriteAllBytes($certFile, $decodedCertData)
	Import-Certificate -CertStoreLocation Cert:\\LocalMachine\Root -FilePath $certFile
	Remove-Item $certFile
	`, encodedCertData)
	addCertScript = base64.StdEncoding.EncodeToString([]byte(addCertScript))

	// workaround for https://github.com/Microsoft/hcsshim/issues/155
	fmt.Printf("%s\n", "Deleting existing containers")
	err = exec.Command(fmt.Sprintf("Get-ComputeProcess | foreach { & %s delete $_.Id }", wincBin)).Run()
	if err != nil {
		return fmt.Errorf("Cannot delete existing containers\n")
	}

	files, err := ioutil.ReadDir(fmt.Sprintf("%s\\volumes", grootDriverStore))
	if !os.IsNotExist(err) {
		return fmt.Errorf("groot delete failed: %s\n", err)
	}

	for _, file := range files {
		err = exec.Command(fmt.Sprintf("%s --driver-store %s delete %s", grootBin, grootDriverStore, file.Name())).Run()
		if err != nil {
			return fmt.Errorf("groot delete failed: %s\n", err)
		}
	}

	fmt.Printf("%s\n", "Begin exporting layer")
	for _, uri := range grootImageUris {
		containerId := fmt.Sprintf("layer%d", int32(time.Now().Unix()))

		fmt.Printf("%s\n", "Creating Volume")
		cmd := exec.Command(fmt.Sprintf("%s --driver-store %s create %s", grootBin, grootDriverStore, uri))
		var stdoutBuffer bytes.Buffer
		cmd.Stdout = &stdoutBuffer
		cmd.Run()
		if err != nil {
			return fmt.Errorf("Groot create failed\n")
		}

		var config map[string]interface{}
		if err := json.Unmarshal(stdoutBuffer.Bytes(), &config); err != nil {
			return fmt.Errorf("failed to parse process spec\n")
		}

		thing := make(map[string]interface{})
		thing["args"] = []string{"powershell.exe", "-EncodedCommand", "$encodedScript"}
		thing["cwd"] = "C:\\"
		config["process"] = thing

		fmt.Printf("Writing config.json")
		bundleDir := filepath.Join(os.TempDir(), containerId)
		configPath := filepath.Join(bundleDir, "config.json")
		if err = os.Mkdir(bundleDir, 0755); err != nil {
			return fmt.Errorf("Failed to create bundle directory\n")
		}

		configBytes, err := json.Marshal(config)
		if err != nil {
			return fmt.Errorf("Failed to write config.json\n")
		}
		configFile, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("Failed to create config.json\n")
		}
		defer configFile.Close()

		_, err = configFile.Write(configBytes)
		if err != nil {
			return fmt.Errorf("Failed to write config.json\n")
		}

		configFile.Sync()

		fmt.Printf("%s\n", "winc run")
		err = exec.Command(fmt.Sprintf("%s run -b %s %s", wincBin, bundleDir, containerId)).Run()
		if err != nil {
			return fmt.Errorf("winc run failed\n")
		}

		fmt.Printf("%s\n", "Running diff-exporter")
		diffOutputFile := filepath.Join(os.TempDir(), fmt.Sprintf("diff-output%d", int32(time.Now().Unix())))
		err = exec.Command(fmt.Sprintf("%s -outputFile %s -containerId %s -bundlePath %s", diffExporterBin, diffOutputFile, containerId, bundleDir)).Run()
		if err != nil {
			return fmt.Errorf("diff-exporter failed\n")
		}

		fmt.Printf("%s\n", "Running hydrator")
		ociImage := strings.Split(uri, "///")[1]
		ociImage = strings.Replace(ociImage, "/", "\\", -1)
		err = exec.Command(fmt.Sprintf("%s add-layer -ociImage %s -layer %s", hydrateBin, ociImage, diffOutputFile)).Run()
		if err != nil {
			return fmt.Errorf("hydrator failed\n")
		}

		fmt.Printf("%s\n", "Cleaning up")
		err = exec.Command(fmt.Sprintf("%s --driver-store %s delete %s", grootBin, grootDriverStore, containerId)).Run()
		if err != nil {
			return fmt.Errorf("groot delete failed\n")
		}
		err = os.RemoveAll(diffOutputFile)
		if err != nil {
			return fmt.Errorf("diff output file deletion failed\n")
		}
	}
	*/
	return nil
}

func main() {
	logger := log.New(os.Stderr, "", 0)
	util := &Util{}
	if err := Run(os.Args, util); err != nil {
		logger.Print(err)
		os.Exit(1)
	}
}