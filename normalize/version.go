package normalize

import (
	"fmt"
	"strconv"

	"github.com/remerge/go-tools/version"
)

// Version returns a noramlized version string in the format major.minor.patch
// if minor and patch are zero only the major verison is returned
func Version(versionStr string) string {
	v := version.Parse(versionStr)

	if v.Major == 0 && v.Minor == 0 && v.Patch == 0 {
		return ""
	}
	if v.Minor == 0 && v.Patch == 0 {
		return strconv.Itoa(v.Major)
	}
	if v.Minor != 0 && v.Patch == 0 {
		return fmt.Sprintf("%d.%d", v.Major, v.Minor)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
