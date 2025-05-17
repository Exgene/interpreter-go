package tokenizer

import (
	"fmt"
	"unicode"
)

var keywords = map[string]TokenType{
	"let":   Let,
	"const": Const,
}

func NewTokenizer(source string) *Tokenizer {
	return &Tokenizer{
		src_code: source,
		idx:      0,
		tokens:   make([]Token, 0),
		buf:      "",
	}
}

func (t TokenType) String() string {
	switch t {
	case Numeric:
		return "Numeric"
	case Identifier:
		return "Identifier"
	case Equals:
		return "Equals"
	case OpenParen:
		return "OpenParen"
	case ClosedParen:
		return "ClosedParen"
	case BinaryOperator:
		return "BinaryOperator"
	case Let:
		return "Let"
	case EOF:
		return "EOF"
	default:
		return fmt.Sprintf("Unknown(%d)", int(t))
	}
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %v, Value: %v}", t.Kind, t.Value)
}

func is_skippable(char rune) bool {
	return unicode.IsSpace(char)
}

func isKeyword(word string) (TokenType, bool) {
	keyword, exists := keywords[word]
	return keyword, exists
}

func Tokenize(code string) {
	tokenizer := NewTokenizer(code)
	tokens := tokenizer.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token.String())
	}
}
