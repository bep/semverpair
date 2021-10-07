package semverpair

import "fmt"

// EncodeSemverPair encodes two semver versions that share the same major version
// into one.
// The encoded version can be decoded back to the originals using DecodeSemver
// as long as minor2 and patch2 is less than 100.
func EncodeSemverPair(major1, minor1, patch1, minor2, patch2 int) (int, int, int) {
	return major1, (minor1 * 100) + minor2, (patch1 * 100) + patch2
}

// DecodeSemver decodes the given semver triplet into the original
// version pair.
// This only works if the original pair's second version's minor and patch version
// was less than 100.
func DecodeSemver(major, minor, patch int) (int, int, int, int, int) {
	minor1 := (minor / 100)
	patch1 := (patch / 100)
	minor2 := minor - (minor1 * 100)
	patch2 := patch - (patch1 * 100)
	return major, minor1, patch1, minor2, patch2
}

// DecodeFirstSemver utility function to decode the first triplet, e.g.
// the upstream version.
func DecodeFirstSemver(major, minor, patch int) (int, int, int) {
	major, minor, patch, _, _ = DecodeSemver(major, minor, patch)
	return major, minor, patch
}

// SemverString is just a helper that formats the version triplet as
// a semver string, e.g. v3.5.2.
func SemverString(major, minor, patch int) string {
	return fmt.Sprintf("v%d.%d.%d", major, minor, patch)
}
