package fqdn

import (
	"bytes"
	"os/exec"
)

// Get uses /bin/hostname to get the fqdn for the host
func Get() string {
	// #nosec
	out, err := exec.Command("/bin/hostname", "-f").Output()
	if err != nil {
		return "localhost"
	}
	return string(bytes.TrimSpace(out))
}
