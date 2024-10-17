package metrics

import (
	"strings"
	"unicode"
)

// SanitizeLabel sanitizes a label to conform with Prometheus label rules
func SanitizeLabel(input string) string {
	input = strings.TrimSpace(input)

	if input == "" {
		return input
	}

	var result strings.Builder

	runeInput := []rune(input)
	// To prevent leading non-letter symbol
	for _, firstRune := range runeInput {
		isAllowedLeadingCharacter := unicode.IsLetter(firstRune) || firstRune == '_'
		if !isAllowedLeadingCharacter {
			// Prepend `_` to a leading digit
			if unicode.IsDigit(firstRune) {
				result.WriteRune('_')
			} else {
				result.WriteRune('_')
				runeInput = runeInput[1:]
			}
		}

		break
	}

	for _, ch := range runeInput {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			result.WriteRune(ch)
		} else {
			result.WriteRune('_')
		}
	}

	// Leading `__` (double undescore) are reserverd for internal use
	sanitized := result.String()
	for strings.HasPrefix(sanitized, "__") {
		sanitized = "_" + sanitized[2:]
	}

	return sanitized
}
