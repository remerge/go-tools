package version

import (
	"regexp"
	"strconv"
	"strings"
)

var versionRegExp = regexp.MustCompile(`(?P<major>\d+)\.?(?P<minor>\d+)?\.?(?P<patch>\d+)?`)

// Version represents a generic version number (For example IOS or Android version)
type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion parses a version string to a Version struct
func Parse(str string) *Version {
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

// Compare a version to another. Will return a value < 0 if the version is smaller, 0 if both are equal and a value > 0 if it is larger
func (version *Version) Compare(other *Version) int {
	cmpMajor := version.Major - other.Major

	if cmpMajor != 0 {
		return cmpMajor
	}

	cmpMinor := version.Minor - other.Minor

	if cmpMinor != 0 {
		return cmpMinor
	}

	return version.Patch - other.Patch
}
