package normalize

// Language tries to normalize given string to ISO-639-1
// or returns empty string
func Language(in string) (out string) {
	if len(in) < 2 || in == "undefined" {
		return ""
	}
	raw := make([]byte, 2)
	var ok bool
	if raw[0], ok = langByte(in[0]); !ok {
		return ""
	}
	if raw[1], ok = langByte(in[1]); !ok {
		return ""
	}
	return string(raw)
}

func langByte(in byte) (out byte, ok bool) {
	if in > 96 && in < 123 {
		// ok, lowercase
		return in, true
	}
	if in > 64 && in < 91 {
		return in+32, true
	}
	return 0, false
}
