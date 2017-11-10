package normalize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var versionRegExp = regexp.MustCompile(`(?P<major>\d+)\.?(?P<minor>\d+)?\.?(?P<patch>\d+)?`)

// Version returns a noramlized version string in the format major.minor.patch
// if minor and patch are zero only the major verison is returned
func Version(version string) string {
	major, minor, patch := versionN(version)

	if major == 0 && minor == 0 && patch == 0 {
		return ""
	}
	if minor == 0 && patch == 0 {
		return strconv.Itoa(major)
	}
	if minor != 0 && patch == 0 {
		return fmt.Sprintf("%d.%d", major, minor)
	}
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func versionN(version string) (major, minor, patch int) {
	if version == "" {
		return 0, 0, 0
	}
	version = strings.Replace(version, "_", ".", -1)
	match := versionRegExp.FindStringSubmatch(version)
	if len(match) == 0 {
		return 0, 0, 0
	}
	// fails for an empty string but that is ok
	major, _ = strconv.Atoi(match[1])
	minor, _ = strconv.Atoi(match[2])
	patch, _ = strconv.Atoi(match[3])
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
	return major, minor, patch
}
