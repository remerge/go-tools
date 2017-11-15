package version

import (
	"fmt"
	"reflect"
	"testing"
)

func inspectVersion(v *Version) string {
	return fmt.Sprintf("Version { Major: %d, Minor: %d, Patch: %d }", v.Major, v.Minor, v.Patch)
}

func TestParse(t *testing.T) {
	testCases := map[string]*Version{
		"11":   &Version{Major: 11},
		"11.1": &Version{Major: 11, Minor: 1},
		"android : 5.0 : L : sdk=21 : LOLLIPOP : sdk=21":    &Version{Major: 5},
		"5.00000000000000000000000000000000000000000000537": &Version{Major: 5, Minor: 537},
		"7.0.0.0.0":             &Version{Major: 7},
		"7.0.1":                 &Version{Major: 7, Patch: 1},
		"7.1.1":                 &Version{Major: 7, Minor: 1, Patch: 1},
		"hey 1.2.3.4.5.6.7.8.9": &Version{Major: 1, Minor: 2, Patch: 3},
		"hey 12_1 3.4.5":        &Version{Major: 12, Minor: 1},
		"05.04":                 &Version{Major: 5, Minor: 4},
		"07.00":                 &Version{Major: 7},
		"iOS 11.1":              &Version{Major: 11, Minor: 1},
		"TEST 6_2_0":            &Version{Major: 6, Minor: 2},
		"20003":                 &Version{Major: 20003},
		"999999999999999.9999999999":                                                                                                                &Version{Major: 1000000, Minor: 1000000},
		"9999999999999999999999999999999999999999.99999999999999999999999":                                                                          &Version{Major: 1000000, Minor: 1000000},
		"999999999999999999999999999999999999999999999999999999999999999999999.9999999999999999999999999999999999999999999999999999999999999999999": &Version{Major: 1000000, Minor: 1000000},
		"0":            &Version{},
		"0.0.0.0 adsf": &Version{},
		"asdasd":       &Version{},
		"null":         &Version{},
		"unknown":      &Version{},
		"other":        &Version{},
		".":            &Version{},
		"_":            &Version{},
	}

	for versionString, expected := range testCases {
		version := Parse(versionString)

		if !reflect.DeepEqual(version, expected) {
			t.Errorf("expect '%s' to be equal to '%s'", inspectVersion(version), inspectVersion(expected))
		}
	}
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
