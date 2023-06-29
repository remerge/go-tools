package fqdn

import (
	"bytes"
	"os"
	"os/exec"
)

// Get returns the FQDN of the host
func Get() string {
	host, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		// #nosec
		out, err := exec.Command("/bin/hostname", "-f").Output()
		if err != nil {
			return "localhost"
		}
		return string(bytes.TrimSpace(out))
	}
	return host
}
