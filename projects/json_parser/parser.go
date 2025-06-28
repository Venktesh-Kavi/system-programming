package json_parser

import (
	"fmt"
	"maps"
)

// Parse the tokens to produce an object
func Parse(tokens []Token) (any, error) {
	token := tokens[0]

	if token.kind != JsonSyntax {
		return nil, UnExpectedTokenError(token, "json type should start with a json syntax, \"{\" or \"[\"")
	}

	json := any(nil)
	var err error
	switch token.value {
	case "{":
		json, err = parseObject(tokens[1:])
		if err != nil {
			return nil, err
		}
		return json, nil
	case "[":
		json, err = parseArray(tokens[1:])
		if err != nil {
			return nil, err
		}
		return json, nil
	default:
		return nil, UnExpectedTokenError(token, "unknown character found, should be a json syntax")
	}
}

func parseObject(tokens []Token) (map[string]any, error) {
	var obj = make(map[string]any)
	if tokens[0].kind != JsonString {
		return nil, UnExpectedTokenError(tokens[0], "unexpected type found, should begin with a json string")
	}

	if tokens[1].kind != JsonSyntax && tokens[1].value != ":" {
		return nil, UnExpectedTokenError(tokens[1], "unexpected type found, should be a json syntax \":\"")
	}

	if tokens[2].kind == JsonSyntax && tokens[2].value == "{" {
		rcxMap, err := parseObject(tokens[3:])
		if err != nil {
			return nil, UnExpectedTokenError(tokens[3], fmt.Sprintf("error parsing object: %v", err))
		}
		obj[tokens[0].value] = rcxMap
		return obj, nil
	} else if tokens[2].kind == JsonSyntax && tokens[2].value == "[" {
		rcxMap, err := parseArray(tokens[3:])
		if err != nil {
			return nil, UnExpectedTokenError(tokens[3], fmt.Sprintf("error parsing array: %v", err))
		}
		obj[tokens[0].value] = rcxMap
		return obj, nil
	} else if !isValueKind(tokens[2].kind) {
		return nil, UnExpectedTokenError(tokens[2], "unexpected value type found, should be string, number, boolean or null")
	}

	var isMultiLine bool
	if tokens[3].kind == JsonSyntax && tokens[3].value == "," {
		// add existing line item to map
		obj[tokens[0].value] = tokens[2].value
		rcxMap, err := parseObject(tokens[4:])
		if err != nil {
			return nil, UnExpectedTokenError(tokens[3], fmt.Sprintf("error parsing object: %v", err))
		}
		maps.Copy(obj, rcxMap)
		isMultiLine = true
	}
	if !isMultiLine && tokens[3].kind != JsonSyntax && tokens[3].value != "}" {
		return nil, UnExpectedTokenError(tokens[3], fmt.Sprintf("unexpected json syntax with closing braces!"))
	} else {
		obj[tokens[0].value] = tokens[2].value
		return obj, nil
	}
}

func isValueKind(kind TokenKind) bool {
	return kind == JsonString || kind == JsonNumber || kind == JsonNull || kind == JsonBoolean
}

func parseArray(tokens []Token) (map[string]any, error) {
	return nil, nil
}
