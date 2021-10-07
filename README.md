
This library speficies a simple **encoding and decoding scheme** to **combine two semver versions** into one.

My intended use for this is as a versioning scheme for Hugo Modules that are wrappers of upstream libraries.

Features:

* The upstream library controls the major version.
* It should be easy to determine the upstream version just by looking at it.
* It should be possible to programmatially decode back to the originals in all practical scenarios.
* A minor or patch increment from either of the two versions should result in a version that is _greater than_ the previous version.

