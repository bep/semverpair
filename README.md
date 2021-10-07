
This library speficies a simple **encoding and decoding scheme** to combine **two semver versions** into one.

My intended use for this is as a versioning scheme for Hugo Modules that are wrappers of upstream libraries. The wrapper and library have diverging release cycles, and we want to maintain the upstream version info, so we cannot just use the upstream version as is.

Features:

* The upstream library controls the major version.
* It should be easy to determine the upstream version just by looking at it.
* It should be possible to programmatially decode back to the originals in all practical scenarios.
* A minor or patch increment from either of the two versions should result in a version that is _greater than_ the previous version.

Have a look at the test code in this repository for some examples. Also see [this module](https://github.com/gohugoio/hugo-mod-jslibs-dist/releases/tag/alpinejs%2Fv3.400.200); here, the upstream version is `v3.4.2` and the combined version `v3.400.200`. If the wrapper needs to make configuration changes it may release a new version, incrementing either the minor or patch version, e.g.: `v3.400.201`.
