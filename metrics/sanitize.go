package metrics

import (
	"strings"
	"unicode"
)

func SanitizeLabel(value string) string {
	var builder strings.Builder
	previousUnderscore := false

	for _, char := range value {
		if char == ' ' || char == '-' {
			if !previousUnderscore {
				builder.WriteRune('_')
				previousUnderscore = true
			}
		} else if unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_' {
			builder.WriteRune(char)
			previousUnderscore = false
		} else {
			if !previousUnderscore {
				builder.WriteRune('_')
				previousUnderscore = true
			}
		}
	}

	// Remove trailing underscore, if any
	result := builder.String()
	return strings.TrimRight(result, "_")
}
