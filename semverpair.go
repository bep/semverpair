package semverpair

import (
	"fmt"
	"strconv"
)

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

// This makes sure that all minor and patch versions start out with an unsigned int (9)
// and that most of them will be four digitis long
const addend = 9000

// Encode encodes two semver versions that share the same major version
// into one.
// The encoded version can be decoded back to the originals using DecodeSemver
// as long as minor2 and patch2 is less than 100.
func Encode(pair Pair) Version {
	minorStr := fmt.Sprintf("9%02d%02d", pair.First.Minor, pair.Second.Minor)
	patchStr := fmt.Sprintf("9%02d%02d", pair.First.Patch, pair.Second.Patch)
	minor := mustAtoi(minorStr)
	patch := mustAtoi(patchStr)

	// v3.1.0", "v3.0.0", "v3.9100.9000
	return Version{
		Major: pair.First.Major,
		Minor: minor,
		Patch: patch,
	}
}

// Decode decodes the given semver triplet into the original
// version pair.
func Decode(v Version) Pair {
	minorStr := strconv.Itoa(v.Minor)
	patchStr := strconv.Itoa(v.Patch)

	if len(minorStr) != 5 {
		panic("minor version is not 4 digits long")
	}

	if minorStr[0] != '9' {
		panic("minor version must start with 9")
	}

	if len(patchStr) != 5 {
		panic("patch version is not 4 digits long")
	}

	if patchStr[0] != '9' {
		panic("patch version must start with 9")
	}

	return Pair{
		First: Version{
			Major: v.Major,
			Minor: mustAtoi(minorStr[1:3]),
			Patch: mustAtoi(patchStr[1:3]),
		},
		Second: Version{
			Major: v.Major,
			Minor: mustAtoi(minorStr[3:]),
			Patch: mustAtoi(patchStr[3:]),
		},
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
