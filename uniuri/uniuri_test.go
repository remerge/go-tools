// Written by Dmitry Chestnykh
// Adapted to use math/rand/v2

package uniuri

import (
	"testing"
)

func TestNew(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := New()
		// Check length
		if len(u) != StdLen {
			t.Fatalf("expected length %d, got %d", StdLen, len(u))
		}
		// Check that it contains only standard characters
		for _, c := range u {
			found := false
			for _, sc := range StdChars {
				if byte(c) == sc {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("unexpected character %c in %q", c, u)
			}
		}
	}
}

func TestNewLen(t *testing.T) {
	for i := 0; i < 100; i++ {
		expectedLen := i + 1
		u := NewLen(expectedLen)
		if len(u) != expectedLen {
			t.Fatalf("expected length %d, got %d", expectedLen, len(u))
		}
		// Check that it contains only standard characters
		for _, c := range u {
			found := false
			for _, sc := range StdChars {
				if byte(c) == sc {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("unexpected character %c in %q", c, u)
			}
		}
	}
}

func TestNewLenChars(t *testing.T) {
	length := 10
	chars := []byte("01234567")
	for i := 0; i < 100; i++ {
		u := NewLenChars(length, chars)
		if len(u) != length {
			t.Fatalf("expected length %d, got %d", length, len(u))
		}
		// Check that it contains only specified characters
		for _, c := range u {
			found := false
			for _, sc := range chars {
				if byte(c) == sc {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("unexpected character %c in %q", c, u)
			}
		}
	}
}

func TestNewLenCharsZeroLength(t *testing.T) {
	u := NewLenChars(0, StdChars)
	if u != "" {
		t.Fatalf("expected empty string, got %q", u)
	}
}

func TestNewLenCharsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for bad charset length")
		}
	}()
	NewLenChars(1, []byte("a"))
}

func TestNewLenCharsPanic256(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for charset length > 256")
		}
	}()
	chars := make([]byte, 257)
	NewLenChars(1, chars)
}

func TestUUIDLen(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := NewLen(UUIDLen)
		if len(u) != UUIDLen {
			t.Fatalf("expected length %d, got %d", UUIDLen, len(u))
		}
	}
}

func TestRandomness(t *testing.T) {
	// Test that two generated strings are different
	// This could theoretically fail but with ~95 bits of entropy the probability is negligible
	u1 := New()
	u2 := New()
	if u1 == u2 {
		t.Fatalf("expected different strings, got %q and %q", u1, u2)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkNewLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewLen(32)
	}
}

func BenchmarkNewLenChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewLenChars(32, StdChars)
	}
}
