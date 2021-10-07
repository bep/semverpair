package semverpair

import "fmt"

// Version represents a semver version.
type Version struct {
	Major int
	Minor int
	Patch int
}

// Pairs represents a pair of semver version sharing the same major version.
type Pair struct {
	First  Version
	Second Version
}

// String formats the version triplet as a semver string, e.g. v3.5.2.
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Encode encodes two semver versions that share the same major version
// into one.
// The encoded version can be decoded back to the originals using DecodeSemver
// as long as minor2 and patch2 is less than 100.
func Encode(pair Pair) Version {
	return Version{
		Major: pair.First.Major,
		Minor: (pair.First.Minor * 100) + pair.Second.Minor,
		Patch: (pair.First.Patch * 100) + pair.Second.Patch,
	}
}

// Decode decodes the given semver triplet into the original
// version pair.
// This only works if the original pair's second version's minor and patch version
// was less than 100.
func Decode(v Version) Pair {
	minor1 := (v.Minor / 100)
	patch1 := (v.Patch / 100)
	minor2 := v.Minor - (minor1 * 100)
	patch2 := v.Patch - (patch1 * 100)

	return Pair{
		First: Version{
			Major: v.Major,
			Minor: minor1,
			Patch: patch1,
		},
		Second: Version{
			Major: v.Major,
			Minor: minor2,
			Patch: patch2,
		},
	}
}
