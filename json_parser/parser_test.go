package json_parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		hasErr   bool
		input    string
		expected map[string]any
	}{
		{
			name:   "test parsing",
			hasErr: false,
			input:  `{"foo": "bar"}`,
			expected: map[string]any{
				"foo": "bar",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := Lex(tc.input)
			if tc.hasErr && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			got, err := Parse(tokens)

			if tc.expected != got {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
