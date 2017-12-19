package normalize

import "github.com/remerge/go-tools/version"

// Version returns a noramlized version string in the format major.minor.patch
// if minor and patch are zero only the major verison is returned
func Version(s []byte) string {
	return version.Parse(s).String()
}
