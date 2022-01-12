
This library speficies a simple **encoding and decoding scheme** to combine **two semver versions** into one.

My intended use for this is as a versioning scheme for Hugo Modules that are wrappers of upstream libraries. The wrapper and library have diverging release cycles, and we want to maintain the upstream version info, so we cannot just use the upstream version as is.

Features:

* The upstream library controls the major version.
* It should be easy to determine the upstream version just by looking at it.
* It should be possible to programmatially decode back to the originals in all practical scenarios.
* A minor or patch increment from either of the two versions should result in a version that is _greater than_ the previous version.

The minor and patch version is represented by the digit 9 and then 4 digits, 2 digits for each version, e.g.:

* `v3.0.0 + v3.0.0 = v3.90000.90000`
* `v3.1.0 + v3.0.0" = v3.90100.90000`
* `v3.6.5 + v3.99.99 = v3.90699.90599`

