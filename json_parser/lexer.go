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

		// Ignore white spaces
		if unicode.IsSpace(runes[0]) {
			if runes[0] == '\n' {
				lineNo++
			}
			runes = runes[1:] // Update the runes slice
			colNo++
			continue
		}

		token, newRunes, err := lexString(runes, lineNo, colNo) // Use newRunes for the result
		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			runes = newRunes // Update the outer runes
			colNo += len(token.value)
		}

		if _, ok := JsonSyntaxChars[string(runes[0])]; ok {
			
		}
		break
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

	escaped := false
	for i, char := range runes[1:] {
		if escaped {
			switch char {
			case 'c', 'b', 'f', 'n', 'r', 't', '\\', '"':
				escaped = false
			default:
				return Token{}, runes, fmt.Errorf("invalid escaped character '%c', at lineNo: %d, colNo: %d", char, lineNo, colNo)
			}
		} else if char == '\\' {
			escaped = true
		} else if char == '"' {
			return Token{
				kind:   JsonSyntax,
				value:  string(runes[1 : i+1]),
				lineNo: lineNo,
				colNo:  colNo,
			}, runes[i:], nil
		}
	}

	// this can happen  only when ending double quote is not found.
	// TODO: handle very long inputs without ending double quotes.
	return Token{}, runes, errors.New("ending double quote not found")
}
