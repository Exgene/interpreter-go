package tokenizer

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	Numeric TokenType = iota
	Identifier
	Equals
	OpenParen
	ClosedParen
	BinaryOperator
	Let
	EOF
)

var keywords = map[string]TokenType{
	"let": Let,
}

type Token struct {
	kind  TokenType
	value *float32
}

type Tokenizer struct {
	idx      int
	buf      string
	tokens   []Token
	src_code string
}

func peek(tokenizer *Tokenizer, ahead int) (string, error) {
	if tokenizer.idx+ahead > len(tokenizer.src_code) {
		return "", fmt.Errorf("Cant peek ahead of string of size (%d): with ahead (%d)", len(tokenizer.src_code), ahead)
	}
	return string(tokenizer.src_code[tokenizer.idx+ahead-1]), nil
}

func NewTokenizer(source string) *Tokenizer {
	return &Tokenizer{
		src_code: source,
		idx:      0,
		tokens:   make([]Token, 0),
		buf:      "",
	}
}

func (t *Tokenizer) isAtEnd() bool {
	if t.idx >= len(t.src_code) {
		return false
	}
	return true
}

func (t *Tokenizer) scanToken() {
	c := t.next()

	switch {
	case c == '(':
		t.addToTokensArray(Token{kind: TokenType(OpenParen), value: nil})
	case c == ')':
		t.addToTokensArray(Token{kind: TokenType(ClosedParen), value: nil})
	case c == '+'|'-'|'/'|'*'|'%':
		t.addToTokensArray(Token{kind: TokenType(BinaryOperator), value: nil})
	case c == '=':
		t.addToTokensArray(Token{kind: TokenType(Equals), value: nil})
	case unicode.IsNumber(c):
		t.scanNumber()
	case unicode.IsLetter(c):
		t.scanLiteral()
	}

}

func (t *Tokenizer) addToTokensArray(token Token) {
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) scanLiteral() {

}

func (t *Tokenizer) scanNumber() {

}

func (t *Tokenizer) next() rune {
	if t.isAtEnd() {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(t.src_code[t.idx:])
	t.idx += 1
	return r
}

func ScanTokens(t *Tokenizer) []Token {
	for !t.isAtEnd() {
		t.scanToken()
	}

	t.tokens = append(t.tokens, Token{TokenType(EOF), nil})
	return t.tokens
}

func is_skippable(char rune) bool {
	return unicode.IsSpace(char)
}

func Tokenize(tokenizer Tokenizer, code string) {
}
