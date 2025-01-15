Learnings

- Learnt usage of typed constants with TokenKind.

```
type TokenKind int

const (
    JsonSyntax TokenKind = iota
    JsonString
)
```

- Bytes are by default for ASCII codes, using them for unicode can result in incorrect manipulation of multi byte
  character (eg.., smiley).
- Learnt how to copy a map into another map. maps.Copy(src, dest)
- Slicing tokens[1:], here 1 is inclusive, tokens[:2] 2 is exclusive.
- In unit testing comparing map[string]any want to return type of any, could not be achieved by just doing got !=
  expected. Manually comparison has to be written, instead starting using testify/assert library.


Lexing

- Initially implementation of lexer was difficult, as I had questions of how to handle the numerous scenarios in
  tokenising the input string.
    - The lexing code is structured as a set of functions were the input strings is passed and seen if it can tokenised.
      Eg.., input string is passed through LexString, LexBoolean and LexNumber function etc.,
- The while loop runes till all runes are analysed and for each iteration of a slice rune, tokenization is done and the
  the rune is re-sliced, by re-assigning to the pre-declared rune.

Parsing

- Parsing involves passing of the tokens and analysing their syntactic correctness.
- Some implementations online use an iterative way to solve this, having checks
  like https://github.com/biraj21/json-parser/blob/main/parser.go
    - The user gives the items to check an iota value and while loops over the length of tokens. He starts of with
      checkKey as initial value of check.
    - The switch case goes through different check conditions and he keeps re-slicing the tokens.
- My approach is that a single unit of json line has fixed syntactic rules and they have designated positions
    - token[0] should be a json string
    - token[1] should be a colon
    - token[2] should be a value
        - value can be nested object/array or a simple json type like (number, string, boolean or null)
    - token[3] can be ',' or }

## TODO
[ ] Lexing Json Numbers
[ ] Parsing Json Arrays

## References

* https://notes.eatonphil.com/writing-a-simple-json-parser.html