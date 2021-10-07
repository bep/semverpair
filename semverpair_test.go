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
		{"v3.6.5", "v3.3.2", "v3.603.502"},
		{"v3.88.43", "v3.3.2", "v3.8803.4302"},
		{"v3.6.5", "v3.13.24", "v3.613.524"},
		{"v3.6.5", "v3.99.99", "v3.699.599"},
		{"v3.88.43", "v3.13.24", "v3.8813.4324"},
	} {
		v1, v2 := splitSemver(test.v1), splitSemver(test.v2)
		encoded := Encode(Pair{First: v1, Second: v2})
		c.Assert(encoded.String(), qt.Equals, test.expect)
		c.Assert(semver.Compare(test.v1, encoded.String()), qt.Equals, -1)
		decoded := Decode(encoded)
		c.Assert(decoded.First.String(), qt.Equals, test.v1)
		c.Assert(decoded.Second, qt.Equals, v2)

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
		{"v3.632.5", "v3.3.2", "v3.63203.502", false},
		{"v3.6.5", "v3.100.2", "v3.700.502", true},
		{"v3.6.5", "v3.3.100", "v3.603.600", true},
		{"v3.1234.5", "3.3.2", "v3.123403.502", false},
		{"v3.6.12345", "v3.3.2", "v3.603.1234502", false},
	} {
		v1, v2 := splitSemver(test.v1), splitSemver(test.v2)
		encoded := Encode(Pair{First: v1, Second: v2})
		c.Assert(encoded.String(), qt.Equals, test.expect)
		c.Assert(semver.Compare(test.v1, encoded.String()), qt.Equals, -1)

		checker := qt.Equals
		if test.overflows {
			checker = qt.Not(qt.Equals)
		}

		c.Assert(Decode(encoded).First.String(), checker, test.v1)

	}
}

func splitSemver(ver string) Version {
	ver = strings.TrimPrefix(ver, "v")
	parts := strings.Split(ver, ".")

	return Version{
		Major: mustAtoi(parts[0]),
		Minor: mustAtoi(parts[1]),
		Patch: mustAtoi(parts[2]),
	}

}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
