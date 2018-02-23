package sha1

import (
	// nolint:golint
	. "crypto/sha1"
	"io"
	"os"
)

// SHA1 calculates the sha1 hash of a byte slice
func SHA1(data []byte) []byte {
	hasher := Sum(data)
	return hasher[:]
}

// nolint:golint
func SHA1S(s string) []byte {
	return SHA1([]byte(s))
}

// nolint:golint
func SHA1F(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		// #nosec
		_ = file.Close()
	}()

	hasher := New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
