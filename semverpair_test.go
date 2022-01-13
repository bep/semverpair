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

		{"v3.1.0", "v3.0.0", "v3.110.100"},
		{"v3.1.1", "v3.0.0", "v3.110.110"},
		{"v3.1.1", "v3.1.1", "v3.111.111"},
		{"v3.0.0", "v3.0.0", "v3.100.100"},
		{"v3.0.0", "v3.1.0", "v3.101.100"},
		{"v3.0.0", "v3.11.0", "v3.20011.100"},
		{"v3.6.5", "v3.3.2", "v3.163.152"},
		{"v3.6.5", "v3.13.24", "v3.20613.20524"},
		{"v3.6.5", "v3.99.99", "v3.20699.20599"},
		{"v3.88.43", "v3.13.24", "v3.28813.24324"},
		{"v3.888888.43", "v3.13.24", "v3.6888888000013.24324"},
	} {

		v1, v2 := splitSemver(test.v1), splitSemver(test.v2)
		c.Assert(semver.IsValid(test.v1), qt.Equals, true)
		c.Assert(semver.IsValid(test.v2), qt.Equals, true)
		encoded := Encode(Pair{First: v1, Second: v2})
		c.Assert(semver.IsValid(encoded.String()), qt.Equals, true, qt.Commentf(encoded.String()))
		c.Assert(encoded.String(), qt.Equals, test.expect)
		c.Assert(semver.Compare(test.v1, encoded.String()), qt.Equals, -1, qt.Commentf("%s < %s", test.v1, encoded.String()))
		decoded := Decode(encoded)
		c.Assert(decoded.First.String(), qt.Equals, test.v1, qt.Commentf(encoded.String()))
		c.Assert(decoded.Second, qt.Equals, v2, qt.Commentf(encoded.String()))

		encodeAndCheckGreaterThan := func(pair Pair, lesserVersion string) {
			encoded := Encode(pair)
			c.Assert(semver.Compare(encoded.String(), lesserVersion), qt.Equals, 1, qt.Commentf("%s > %s", encoded.String(), lesserVersion))
		}

		basePair := Pair{First: v1, Second: v2}
		lesserVersion := Encode(basePair).String()
		basePair.First.Patch++
		encodeAndCheckGreaterThan(basePair, lesserVersion)
		basePair.First.Minor++
		encodeAndCheckGreaterThan(basePair, lesserVersion)
		basePair.Second.Patch++
		encodeAndCheckGreaterThan(basePair, lesserVersion)
		basePair.Second.Minor++
		encodeAndCheckGreaterThan(basePair, lesserVersion)

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
