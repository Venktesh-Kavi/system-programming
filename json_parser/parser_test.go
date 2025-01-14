package json_parser

import (
	"github.com/stretchr/testify/assert"
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
			input:  `{"foo": "bar", "one": {"bar": "baz"}}`,
			expected: map[string]any{
				"foo": "bar",
				"one": map[string]any{
					"bar": "baz",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := Lex(tc.input)
			if !tc.hasErr && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			got, err := Parse(tokens)
			assert.Equal(t, tc.expected, got)
		})
	}
}
