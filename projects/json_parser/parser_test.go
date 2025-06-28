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
		errCause string
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
		{
			name:     "error parsing",
			hasErr:   true,
			input:    `{{`,
			expected: map[string]any{},
			errCause: "unexpected token found: JsonSyntax, lineNo: 1, colNo: 2, reason: unexpected type found, should begin with a json string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := Lex(tc.input)
			got, err := Parse(tokens)

			if !tc.hasErr && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if tc.hasErr && err != nil {
				assert.EqualError(t, err, tc.errCause)
				return
			}
			assert.Equal(t, tc.expected, got)
		})
	}
}
