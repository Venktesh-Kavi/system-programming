package json_parser

import "fmt"

type TokenKind int

const (
	JsonBoolean TokenKind = iota
	JsonString
	JsonNumber
	JsonNull
	JsonSyntax
)

const (
	JsonTrue  = "true"
	JsonFalse = "false"
)

var JsonSyntaxChars = map[string]struct{}{
	"{": {},
	"}": {},
	":": {},
	"[": {},
	"]": {},
	",": {},
}

func convertKindToString(kind TokenKind) string {
	switch kind {
	case JsonBoolean:
		return "JsonBoolean"
	case JsonString:
		return "JsonString"
	case JsonNumber:
		return "JsonNumber"
	case JsonSyntax:
		return "JsonSyntax"
	case JsonNull:
		return "JsonNull"
	default:
		return ""
	}
}

type Token struct {
	kind   TokenKind
	value  string
	lineNo int
	colNo  int
}

func UnExpectedTokenError(token Token, msg string) error {
	return fmt.Errorf("unexpected token found: %s, lineNo: %d, colNo: %d, reason: %s", convertKindToString(token.kind), token.lineNo, token.colNo, msg)
}
