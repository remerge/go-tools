package fqdn

import (
	"bytes"
	"os"
	"os/exec"
)

// Get tries to get the hostname from the HOSTNAME environment variable
// and falls back to the output of the hostname command
func Get() string {
	host, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		// #nosec
		out, err := exec.Command("/bin/hostname", "-f").Output()
		if err != nil {
			return "unknown"
		}
		return string(bytes.TrimSpace(out))
	}
	return host
}
