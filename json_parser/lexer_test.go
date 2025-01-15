package json_parser

import (
	"reflect"
	"testing"
)

// tabular test cases
func TestLexing(t *testing.T) {
	tCases := []struct {
		name   string
		input  string
		want   []Token
		hasErr bool
	}{
		{
			name:  "json with empty space as first character",
			input: ` {"foo": "bar"}`,
			want: []Token{
				{kind: JsonSyntax, value: "{", lineNo: 1, colNo: 2},
				{kind: JsonString, value: "foo", lineNo: 1, colNo: 3},
				{kind: JsonSyntax, value: ":", lineNo: 1, colNo: 8},
				{kind: JsonString, value: "bar", lineNo: 1, colNo: 10},
				{kind: JsonSyntax, value: "}", lineNo: 1, colNo: 15},
			},
		},
		{
			name:  "json with boolean value",
			input: `{"foo": true}`,
			want: []Token{
				{kind: JsonSyntax, value: "{", lineNo: 1, colNo: 1},
				{kind: JsonString, value: "foo", lineNo: 1, colNo: 2},
				{kind: JsonSyntax, value: ":", lineNo: 1, colNo: 7},
				{kind: JsonBoolean, value: "true", lineNo: 1, colNo: 9},
				{kind: JsonSyntax, value: "}", lineNo: 1, colNo: 13},
			},
		},
		{
			name:  "error cases",
			input: `{{`,
			want: []Token{
				{kind: JsonSyntax, value: "{", lineNo: 1, colNo: 1},
				{kind: JsonSyntax, value: "{", lineNo: 1, colNo: 2},
			},
			hasErr: true,
		},
	}

	for i, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Lex(tc.input)
			if (err != nil) && !tc.hasErr {
				t.Fatalf("test %d: unexpected error: %v", i, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("test %d:\ngot  %v\nwant %v", i, got, tc.want)
			}
		})
	}
}
