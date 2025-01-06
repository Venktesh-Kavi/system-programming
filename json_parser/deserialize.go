package json_parser

import "fmt"

func Deserialize(input string) (any, error) {
	token, err := Lex(input)
	Tokenize(token)
	if err != nil {
		return nil, err
	}

	//json := Parse(token)
	fmt.Println("Tokens generated successfully:", token)

	return token, nil
}
