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
	firstCh := runeInput[0]
	// To prevent leading non-letter symbol
	isAllowedLeadingCharacter := unicode.IsLetter(firstCh)
	if !isAllowedLeadingCharacter {
		result.WriteRune('_')
	}

	for _, ch := range runeInput {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result.WriteRune(ch)
		} else {
			result.WriteRune('_')
		}
	}

	// Leading `__` (double undescore) are reserverd for internal use
	sanitized := result.String()
	for strings.HasPrefix(sanitized, "__") {
		sanitized = sanitized[1:]
	}

	return sanitized
}
