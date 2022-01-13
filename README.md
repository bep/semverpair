
This library speficies a simple **encoding and decoding scheme** to combine **two semver versions** into one.

My intended use for this is as a versioning scheme for Hugo Modules that are wrappers of upstream libraries. The wrapper and library have diverging release cycles, and we want to maintain the upstream version info, so we cannot just use the upstream version as is.

Features:

* The upstream library controls the major version.
* It should be easy to determine the upstream version just by looking at it.
* It should always be possible to programmatially decode back to the originals.
* A minor or patch increment from either of the two versions should result in a version that is _greater than_ the previous version.

The minor and patch version:

* Starts out with one digit (width) telling how many digits each version holds.
* Then each version is printed padded with leading zeros.

Some examples a width of 2 digits per version:

* `v3.0.0 + v3.0.0 = v3.20000.20000`
* `v3.1.0 + v3.0.0" = v3.20100.20000`
* `v3.6.5 + v3.99.99 = v3.20699.20599`

If either of the versions reaches 100, the prefix can be incremented:

* `v3.6.5 + v3.99.100 = v3.20699.3005100`