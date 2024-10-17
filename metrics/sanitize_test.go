package metrics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizePrometheusLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-label value-with-special@chars!", "my_label_value_with_special_chars_"},
		{"simple-label", "simple_label"},
		{"  label with spaces  ", "label_with_spaces"},
		{"@label with spaces  ", "_label_with_spaces"},
		{" @label with spaces  ", "_label_with_spaces"},
		{"label-with-dashes", "label_with_dashes"},
		{"123-label!", "_123_label_"},
		{"label_with_underscores", "label_with_underscores"},
		{"complex_label- test@1!2", "complex_label__test_1_2"},
		{"a-_label", "a__label"},
		{"", ""},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("case_%d", idx), func(t *testing.T) {
			result := SanitizeLabel(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
