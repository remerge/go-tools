package normalize

// Language tries to normalize given string to ISO-639-1
// or returns empty string
func Language(in string) string {
	inSize := len(in)
	if inSize < 2 || (inSize > 2 && (isByteLowercase(in[2]) || isByteUppercase(in[2]))) {
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
	if isByteLowercase(in) {
		// ok, lowercase
		return in, true
	}
	if isByteUppercase(in) {
		return in+32, true
	}
	return 0, false
}

func isByteLowercase(in byte) bool  {
	return in > 96 && in < 123
}

func isByteUppercase(in byte) bool  {
	return in > 64 && in < 91
}
