package semverpair

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

// Encode encodes two semver versions that share the same major version
// into one.
// The encoded version can be decoded back to the originals using Decode.
func Encode(pair Pair) Version {
	widthMinor := maxWidth(pair.First.Minor, pair.Second.Minor)
	widthPatch := maxWidth(pair.First.Patch, pair.Second.Patch)
	formatMinor := strings.ReplaceAll("x%0xd%0xd", "x", strconv.Itoa(widthMinor))
	formatPatch := strings.ReplaceAll("x%0xd%0xd", "x", strconv.Itoa(widthPatch))
	minorStr := fmt.Sprintf(formatMinor, pair.First.Minor, pair.Second.Minor)
	patchStr := fmt.Sprintf(formatPatch, pair.First.Patch, pair.Second.Patch)
	minor := mustAtoi(minorStr)
	patch := mustAtoi(patchStr)

	return Version{
		Major: pair.First.Major,
		Minor: minor,
		Patch: patch,
	}
}

func maxWidth(n1, n2 int) int {
	w1, w2 := width(n1), width(n2)
	if w1 > w2 {
		return w1
	}
	return w2
}

func width(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n)) + 1)
}

// Decode decodes the given semver triplet into the original
// version pair.
func Decode(v Version) Pair {
	minorStr := strconv.Itoa(v.Minor)
	patchStr := strconv.Itoa(v.Patch)

	widthMinor := mustAtoi(minorStr[:1])
	widthPatch := mustAtoi(patchStr[:1])

	return Pair{
		First: Version{
			Major: v.Major,
			Minor: mustAtoi(minorStr[1 : widthMinor+1]),
			Patch: mustAtoi(patchStr[1 : widthPatch+1]),
		},
		Second: Version{
			Major: v.Major,
			Minor: mustAtoi(minorStr[widthMinor+1:]),
			Patch: mustAtoi(patchStr[widthPatch+1:]),
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
