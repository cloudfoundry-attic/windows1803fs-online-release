# windows1803fs-online-release

A [BOSH](http://docs.cloudfoundry.org/bosh/) release for deploying [windows1803fs](https://github.com/cloudfoundry-incubator/windows2016fs/tree/master/1803).

**Note:**

This release assumes your BOSH installation has internet access at deploy time.

## smoke test

Ensure that `winc-release` and `windows1803fs-release` are uploaded to your BOSH director.

```
bosh -d windows1803fs-smoke-test deploy manifests/smoke-test.yml
```
