$ErrorActionPreference = "Stop";
trap { $host.SetShouldExit(1) }
Write-Host "Starting windows1803fs pre-start"

$certData = "<%= p("windows-rootfs.trusted_certs") %>"
$certFile=[System.IO.Path]::GetTempFileName()
$certData | Out-File $certFile
$grootDriverStore = "<%= p("groot.driver_store") %>"
$grootImageUris = <% p("groot.cached_image_uris").join(" ") %>
$certInjectorBin = "c:\var\vcap\packages\certificate-injector\certificate-injector.exe"
$certInjectorBin $grootDriverStore $certFile $grootImageUris

# Happy path test!

$grootBin = "c:\var\vcap\packages\groot\groot.exe"
$wincBin = "c:\var\vcap\packages\winc\winc.exe"
$containerId = "windows1803fs-smoke" + [System.Math]::Round((date -UFormat %s),0)
$grootImageUri = <% p("groot.cached_image_uris")[0] %>
$stdOut = (& $grootBin --driver-store $grootDriverStore create $grootImageUri $containerId)
$processSpec=@"
{"args": ["powershell.exe", "-EncodedCommand", "$encodedScript"], "cwd": "C:\\" }
"@

$pObj = $processSpec | convertfrom-json
$config | Add-Member -Force -Name "process" -Value $pObj -MemberType NoteProperty

Write-Host "Writing config.json"
$bundleDir = Join-Path $env:TEMP $containerId
$configPath = Join-Path $bundleDir "config.json"
rm -Recurse -Force -ErrorAction SilentlyContinue $bundleDir
mkdir $bundleDir | Out-Null
$configJson = ($config | ConvertTo-Json)
Set-Content -Path $configPath -Value $configJson

Write-Host "winc run"
& $wincBin run -b $bundleDir $containerId
$stdOut = (& $wincBin exec "powershell ls Cert:\\LocalMachine\Root")

$certificateObject = New-Object System.Security.Cryptography.X509Certificates.X509Certificate2
$certificateObject.Import($CertificatePath, $sSecStrPassword, [System.Security.Cryptography.X509Certificates.X509KeyStorageFlags]::DefaultKeySet)
$expectedThumbprint = $certificateObject.Thumbprint
$testSuccess = $stdOut.Contains($expectedThumbprint)

Remove-Item $certFile

if ($testSuccess) {
  echo "Test succeeded"
  exit 0
}
echo "Test failed. Certificate does not exist in the container"
exit 1
