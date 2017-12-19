package version

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func inspectVersion(v *Version) string {
	return fmt.Sprintf("Version { Major: %d, Minor: %d, Patch: %d }", v.Major, v.Minor, v.Patch)
}

func TestParse(t *testing.T) {
	testCases := map[string]*Version{
		"11":   {Major: 11},
		"11.1": {Major: 11, Minor: 1},
		"android : 5.0 : L : sdk=21 : LOLLIPOP : sdk=21":    {Major: 5},
		"5.00000000000000000000000000000000000000000000537": {Major: 5, Minor: 537},
		"7.0.0.0.0":             {Major: 7},
		"7.0.1":                 {Major: 7, Patch: 1},
		"7.1.1":                 {Major: 7, Minor: 1, Patch: 1},
		"hey 1.2.3.4.5.6.7.8.9": {Major: 1, Minor: 2, Patch: 3},
		"hey 12_1 3.4.5":        {Major: 12, Minor: 1},
		"05.04":                 {Major: 5, Minor: 4},
		"07.00":                 {Major: 7},
		"iOS 11.1":              {Major: 11, Minor: 1},
		"TEST 6_2_0":            {Major: 6, Minor: 2},
		"20003":                 {Major: 20003},
		"999999999999999.9999999999":                                                                                                                {Major: 1000000, Minor: 1000000},
		"9999999999999999999999999999999999999999.99999999999999999999999":                                                                          {Major: 1000000, Minor: 1000000},
		"999999999999999999999999999999999999999999999999999999999999999999999.9999999999999999999999999999999999999999999999999999999999999999999": {Major: 1000000, Minor: 1000000},
		"0":            {},
		"0.0.0.0 adsf": {},
		"asdasd":       {},
		"null":         {},
		"unknown":      {},
		"other":        {},
		".":            {},
		"_":            {},
	}

	for versionString, expected := range testCases {
		version := Parse([]byte(versionString))

		if !reflect.DeepEqual(version, expected) {
			fmt.Printf("expect %v => '%s' to be equal to '%s'\n", versionString, inspectVersion(version), inspectVersion(expected))
			t.Errorf("expect %v => '%s' to be equal to '%s'", versionString, inspectVersion(version), inspectVersion(expected))
		}
	}
}

func compareStringOutput(t *testing.T, v *Version, expectedOutput string) {
	if v.String() != expectedOutput {
		t.Errorf("expected %s to be %s", v, expectedOutput)
	}
}

func TestVersionString(t *testing.T) {
	compareStringOutput(t, &Version{}, "")
	compareStringOutput(t, &Version{Major: 1}, "1")
	compareStringOutput(t, &Version{Minor: 2}, "0.2")
	compareStringOutput(t, &Version{Major: 1, Minor: 2}, "1.2")
	compareStringOutput(t, &Version{Patch: 3}, "0.0.3")
	compareStringOutput(t, &Version{Major: 1, Patch: 3}, "1.0.3")
	compareStringOutput(t, &Version{Minor: 2, Patch: 3}, "0.2.3")
	compareStringOutput(t, &Version{Major: 1, Minor: 2, Patch: 3}, "1.2.3")
}

func TestVersionCompare(t *testing.T) {
	version := Version{Major: 5, Minor: 5, Patch: 5}

	if version.Compare(&Version{}) < 1 {
		t.Errorf("expected %s to be larger than 0.0.0", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 4, Minor: 6, Patch: 7}) < 1 {
		t.Errorf("expected %s to be larger than 4.6.7", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5}) < 1 {
		t.Errorf("expected %s to be larger than 5", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 4, Patch: 7}) < 1 {
		t.Errorf("expected %s to be larger than 5.4.7", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 5}) < 1 {
		t.Errorf("expected %s to be larger than 5.4", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 5, Patch: 4}) < 1 {
		t.Errorf("expected %s to be larger than 5.5.4", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 5, Patch: 5}) != 0 {
		t.Errorf("expected %s to be equal to 5.5.5", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 6, Minor: 4, Patch: 3}) > -1 {
		t.Errorf("expected %s to be smaller than 6.4.3", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 6}) > -1 {
		t.Errorf("expected %s to be smaller than 6", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 6}) > -1 {
		t.Errorf("expected %s to be smaller than 5.5", inspectVersion(&version))
	}

	if version.Compare(&Version{Major: 5, Minor: 5, Patch: 6}) > -1 {
		t.Errorf("expected %s to be smaller than 5.5.6", inspectVersion(&version))
	}
}

var versionRegExp = regexp.MustCompile(`(?P<major>\d+)\.?(?P<minor>\d+)?\.?(?P<patch>\d+)?`)

func regexParse(str string) *Version {
	version := &Version{}
	if str == "" {
		return version
	}
	str = strings.Replace(str, "_", ".", -1)
	match := versionRegExp.FindStringSubmatch(str)
	if len(match) == 0 {
		return version
	}
	// fails for an empty string but that is ok
	major, _ := strconv.Atoi(match[1])
	minor, _ := strconv.Atoi(match[2])
	patch, _ := strconv.Atoi(match[3])
	// remove some insane versions
	if major > 1000000 {
		major = 1000000
	}
	if minor > 1000000 {
		minor = 1000000
	}
	if patch > 1000000 {
		patch = 1000000
	}

	version.Major = major
	version.Minor = minor
	version.Patch = patch

	return version
}

func BenchmarkVersionRegex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			regexParse("7.1.1")
		}
	})
}

func BenchmarkVersionRagel(b *testing.B) {
	v := []byte("7.1.1")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Parse(v)
		}
	})
}
