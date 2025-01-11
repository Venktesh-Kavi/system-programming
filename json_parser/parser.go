package json_parser

import (
	"errors"
	"maps"
)

// Parse the tokens to produce an object
func Parse(tokens []Token) (any, error) {
	token := tokens[0]

	if token.kind != JsonSyntax {
		return nil, UnExpectedTokenError(token)
	}

	json := []any{}
	switch token.value {
	case "{":
		obj, err := parseObject(tokens[1:])
		if err != nil {
			return nil, err
		}
		json = append(json, obj)
	case "[":
		parseArray(tokens[1:])
	default:
		return nil, UnExpectedTokenError(token)
	}

	return json, nil
}

func parseObject(tokens []Token) (map[string]any, error) {
	var obj = make(map[string]any)
	if tokens[0].kind != JsonString {
		return nil, UnExpectedTokenError(tokens[0])
	}

	if tokens[1].kind != JsonSyntax && tokens[1].value != ":" {
		return nil, UnExpectedTokenError(tokens[1])
	}

	if !isValueKind(tokens[2].kind) {
		return nil, UnExpectedTokenError(tokens[2])
	}

	if tokens[3].kind == JsonSyntax && tokens[3].value == "," {
		rcxMap, err := parseObject(tokens[4:])
		maps.Copy(obj, rcxMap)
		return obj, err
	} else if tokens[3].kind == JsonSyntax && tokens[3].value == "}" {
		obj[tokens[0].value] = tokens[2].value
		return obj, nil
	}
	return nil, errors.New("illegal state error")
}

func isValueKind(kind TokenKind) bool {
	return kind == JsonString || kind == JsonNumber || kind == JsonNull || kind == JsonBoolean
}

func parseArray(tokens []Token) error {
	return nil
}
