package semverpair

import (
	"strconv"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"golang.org/x/mod/semver"
)

func TestSemverPair(t *testing.T) {
	c := qt.New(t)

	for _, test := range []struct {
		v1     string
		v2     string
		expect string
	}{
		{"v3.6.5", "v0.3.2", "v3.603.502"},
		{"v3.88.43", "v0.3.2", "v3.8803.4302"},
		{"v3.6.5", "v0.13.24", "v3.613.524"},
		{"v3.6.5", "v0.99.99", "v3.699.599"},
		{"v3.88.43", "v0.13.24", "v3.8813.4324"},
	} {
		major, minor1, patch1 := splitSemver(test.v1)
		_, minor2, patch2 := splitSemver(test.v2)
		encMajor, encMinor, encPatch := EncodeSemverPair(major, minor1, patch1, minor2, patch2)
		encoded := SemverString(encMajor, encMinor, encPatch)
		c.Assert(encoded, qt.Equals, test.expect)
		c.Assert(semver.Compare(test.v1, encoded), qt.Equals, -1)
		c.Assert(SemverString(DecodeFirstSemver(encMajor, encMinor, encPatch)), qt.Equals, test.v1)
		_, _, _, minor2Decoded, patch2Decoded := DecodeSemver(encMajor, encMinor, encPatch)
		c.Assert(minor2Decoded, qt.Equals, minor2)
		c.Assert(patch2Decoded, qt.Equals, patch2)

	}
}

func TestSemverPairOverflow(t *testing.T) {
	c := qt.New(t)

	for _, test := range []struct {
		v1        string
		v2        string
		expect    string
		overflows bool
	}{
		{"v3.632.5", "v0.3.2", "v3.63203.502", false},
		{"v3.6.5", "v0.100.2", "v3.700.502", true},
		{"v3.6.5", "v0.3.100", "v3.603.600", true},
		{"v3.1234.5", "v0.3.2", "v3.123403.502", false},
		{"v3.6.12345", "v0.3.2", "v3.603.1234502", false},
	} {
		major, minor1, patch1 := splitSemver(test.v1)
		_, minor2, patch2 := splitSemver(test.v2)
		encMajor, encMinor, encPatch := EncodeSemverPair(major, minor1, patch1, minor2, patch2)
		encoded := SemverString(encMajor, encMinor, encPatch)
		c.Assert(encoded, qt.Equals, test.expect)
		c.Assert(semver.Compare(test.v1, encoded), qt.Equals, -1)

		checker := qt.Equals
		if test.overflows {
			checker = qt.Not(qt.Equals)
		}

		c.Assert(SemverString(DecodeFirstSemver(encMajor, encMinor, encPatch)), checker, test.v1)

	}
}

func splitSemver(ver string) (int, int, int) {
	ver = strings.TrimPrefix(ver, "v")
	parts := strings.Split(ver, ".")

	return mustAtoi(parts[0]), mustAtoi(parts[1]), mustAtoi(parts[2])

}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
