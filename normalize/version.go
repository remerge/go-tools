package normalize

import "github.com/remerge/go-tools/version"

// Version returns a noramlized version string in the format major.minor.patch
// if minor and patch are zero only the major version is returned
func Version(s string) string {
	return version.Parse(s).String()
}
