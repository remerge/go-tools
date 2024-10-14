package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizePrometheusLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-label value-with-special@chars!", "my_label_value_with_special_chars"},
		{"simple-label", "simple_label"},
		{"  label with spaces  ", "_label_with_spaces"},
		{"label-with-dashes", "label_with_dashes"},
		{"123-label!", "123_label"},
		{"label_with_underscores", "label_with_underscores"},
		{"complex_label- test@1!2", "complex_label_test_1_2"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := SanitizeLabel(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
