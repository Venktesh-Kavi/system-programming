package json_parser

import (
	"errors"
	"fmt"
	"unicode"
)

func Lex(input string) ([]Token, error) {
	var tokens []Token
	lineNo, colNo := 1, 1
	runes := []rune(input) // Declare runes outside the loop
	for len(runes) > 0 {
		token := Token{}
		var err error
		if unicode.IsSpace(runes[0]) {
			if runes[0] == '\n' {
				lineNo++
			}
			runes = runes[1:] // Update the runes slice
			colNo++
			continue
		}

		token, runes, err = lexString(runes, lineNo, colNo) // Use newRunes for the result
		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			colNo += len(token.value) + 2
			continue
		}

		if _, ok := JsonSyntaxChars[string(runes[0])]; ok {
			tokens = append(tokens, Token{
				kind:   JsonSyntax,
				value:  string(runes[0]),
				lineNo: lineNo,
				colNo:  colNo,
			})
			runes = runes[1:]
			colNo++
		} else {
			return tokens, fmt.Errorf("unexpected character %s, at lineNo %d and colNo %d", string(runes[0]), lineNo, colNo)
		}
	}
	return tokens, nil
}

func lexNumber(runes []rune, lineNo, colNo int) (Token, error) {
	return Token{}, nil
}

//func lexNumber(runes []rune, lineNo, colNo int) (Token, error) {
//	if !unicode.IsDigit(runes[0]) {
//		return Token{}, nil
//	}
//
//	var endsAt int = len(runes) - 1
//	for i, c := range runes {
//		if !unicode.IsNumber(c) && c != '.' && c != '_' {
//			endsAt = i
//			return Token{}, nil
//		}
//	}
//}

func lexString(runes []rune, lineNo, colNo int) (Token, []rune, error) {
	if runes[0] != '"' {
		return Token{}, runes, nil
	}

	runes = runes[1:]

	escaped := false

	for i, char := range runes {
		if escaped {
			switch char {
			case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
				escaped = false
			default:
				return Token{}, runes, fmt.Errorf("invalid escaped character '\\%s' at line %d, col %d", string(char), lineNo, i+colNo)
			}
		} else if char == '\\' {
			escaped = true
		} else if char == '"' {
			return Token{JsonString, string(runes[:i]), lineNo, colNo}, runes[i+1:], nil
		}
	}

	// this can happen  only when ending double quote is not found.
	return Token{}, runes, errors.New("ending double quote not found")
}