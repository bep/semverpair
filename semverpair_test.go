package semverpair

import (
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

		{"v3.88.43", "v3.3.2", "v3.98803.94302"},

		{"v3.1.0", "v3.0.0", "v3.90100.90000"},
		{"v3.1.1", "v3.0.0", "v3.90100.90100"},
		{"v3.1.1", "v3.1.1", "v3.90101.90101"},
		{"v3.0.0", "v3.0.0", "v3.90000.90000"},
		{"v3.0.0", "v3.1.0", "v3.90001.90000"},
		{"v3.0.0", "v3.11.0", "v3.90011.90000"},
		{"v3.6.5", "v3.3.2", "v3.90603.90502"},
		{"v3.6.5", "v3.13.24", "v3.90613.90524"},
		{"v3.6.5", "v3.99.99", "v3.90699.90599"},
		{"v3.88.43", "v3.13.24", "v3.98813.94324"},
	} {

		v1, v2 := splitSemver(test.v1), splitSemver(test.v2)
		c.Assert(semver.IsValid(test.v1), qt.Equals, true)
		c.Assert(semver.IsValid(test.v2), qt.Equals, true)
		encoded := Encode(Pair{First: v1, Second: v2})
		c.Assert(semver.IsValid(encoded.String()), qt.Equals, true, qt.Commentf(encoded.String()))
		c.Assert(encoded.String(), qt.Equals, test.expect)
		if encoded.String() != test.v1 {
			c.Assert(semver.Compare(test.v1, encoded.String()), qt.Equals, -1, qt.Commentf("%s < %s", test.v1, encoded.String()))
		}
		decoded := Decode(encoded)
		c.Assert(decoded.First.String(), qt.Equals, test.v1, qt.Commentf(encoded.String()))
		c.Assert(decoded.Second, qt.Equals, v2, qt.Commentf(encoded.String()))

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
