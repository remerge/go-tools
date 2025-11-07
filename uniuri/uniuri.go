// Written by Dmitry Chestnykh
// Adapted to use math/rand/v2

// Package uniuri generates random strings good for use in URIs to identify
// unique objects.
//
// Example usage:
//
//	s := uniuri.New() // s is now "apHCJBl7L1OmC57n"
//
// A standard string is 16 characters of [A-Za-z0-9] and is safe for use in URLs.
package uniuri

import (
	"math/rand/v2"
)

const (
	// StdLen is a standard length of uniuri string to achieve ~95 bits of entropy.
	StdLen = 16
	// UUIDLen is a length of uniuri string to achieve ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)

// StdChars is a set of standard characters allowed in uniuri string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// New returns a new random string of the standard length, consisting of
// standard characters.
func New() string {
	return NewLenChars(StdLen, StdChars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func NewLen(length int) string {
	return NewLenChars(length, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // allocate extra space for retries
	i := 0
	for {
		// Fill random byte slice using math/rand/v2
		for j := range r {
			r[j] = byte(rand.Uint32())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
