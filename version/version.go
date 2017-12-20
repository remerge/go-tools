//go:generate ragel -Z -G2 version_fsm.rl
package version

import (
	"fmt"
	"strconv"
)

// Version represents a generic version number (For example IOS or Android
// version)
type Version struct {
	Major int
	Minor int
	Patch int
}

func (version *Version) String() string {
	if version.Major == 0 && version.Minor == 0 && version.Patch == 0 {
		return ""
	}
	if version.Minor == 0 && version.Patch == 0 {
		return strconv.Itoa(version.Major)
	}
	if version.Patch == 0 {
		return fmt.Sprintf("%d.%d", version.Major, version.Minor)
	}
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

// Compare a version to another. Will return a value < 0 if the version is
// smaller, 0 if both are equal and a value > 0 if it is larger
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
