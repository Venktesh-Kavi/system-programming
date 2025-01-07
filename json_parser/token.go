package json_parser

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

type Token struct {
	kind   TokenKind
	value  string
	lineNo int
	colNo  int
}
