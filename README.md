# windows1803fs-online-release

A [BOSH](http://docs.cloudfoundry.org/bosh/) release for deploying [windows1803fs](https://github.com/cloudfoundry-incubator/windows2016fs/tree/master/1803).

**Note:**

This release assumes your BOSH installation has internet access at deploy time.

## smoke test

Ensure that `winc-release` and `windows1803fs-release` are uploaded to your BOSH director.

```
bosh -d windows1803fs deploy manifests/smoke-test.yml \
  -v fs-version="windows1803fs" \
  -v stemcell-os="windows1803" \
  -v smoke-test-name="windows1803fs-smoke-test" \
  -v cached-image-uris=["oci:///C:/var/vcap/packages/windows1803fs"]
bosh -d windows1803fs run-errand smoke-test
```

## Requirements

* This bosh release can only be deployed together with a [winc-release](https://github.com/cloudfoundry/winc-release) of v2.0 or higher. The [windows1803fs pre-start script](/jobs/windows1803fs/templates/pre-start.ps1.erb) waits for winc-release's groot pre-start to signal.
