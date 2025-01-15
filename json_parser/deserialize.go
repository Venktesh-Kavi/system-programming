package json_parser

func Deserialize(input string) (any, error) {
	token, err := Lex(input)
	json, err := Parse(token)
	if err != nil {
		return nil, err
	}
	return json, nil
}
