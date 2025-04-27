package tokenizer

import (
	"fmt"
	"strconv"
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

func (tokenizer *Tokenizer) peek(ahead int) rune {
	if tokenizer.idx+ahead > len(tokenizer.src_code) {
		fmt.Errorf("Shouldnt happen")
	}
	c, _ := utf8.DecodeRuneInString(tokenizer.src_code[tokenizer.idx+ahead-1:])
	return c
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
		t.scanNumber(c)
	case unicode.IsLetter(c):
		t.scanLiteral(c)
	}

}

func (t *Tokenizer) addToTokensArray(token Token) {
	t.tokens = append(t.tokens, token)
}

func isKeyword(word string) (TokenType, bool) {
	keyword, exists := keywords[word]
	return keyword, exists
}

func (t *Tokenizer) scanLiteral(c rune) {
	t.buf += string(c)
	for !t.isAtEnd() && unicode.IsLetter(t.peek(1)) {
		t.buf += string(t.next())
	}

	tokenType, isBufAKeyword := isKeyword(t.buf)

	if isBufAKeyword {
		t.addToTokensArray(Token{kind: tokenType, value: nil})
	} else {
		t.addToTokensArray(Token{kind: TokenType(Identifier), value: nil})
	}

	t.buf = ""
}

func (t *Tokenizer) scanNumber(c rune) {
	t.buf += string(c)
	for !t.isAtEnd() && unicode.IsNumber(t.peek(1)) {
		t.buf += string(t.next())
	}
	value, err := strconv.ParseFloat(t.buf, 32)
	value32 := float32(value)
	if err != nil {
		fmt.Errorf("Error parsing float from buf, (%s)::", t.buf)
	}
	valuePtr := &value32
	t.addToTokensArray(Token{kind: Numeric, value: valuePtr})
	t.buf = ""
}

func (t *Tokenizer) next() rune {
	if t.isAtEnd() {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(t.src_code[t.idx:])
	t.idx += 1
	return r
}

func (t *Tokenizer) ScanTokens() []Token {
	for !t.isAtEnd() {
		t.scanToken()
	}

	t.addToTokensArray(Token{TokenType(EOF), nil})
	return t.tokens
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %d, Lexeme: %f", t.kind, t.value)
}

func is_skippable(char rune) bool {
	return unicode.IsSpace(char)
}

func Tokenize(code string) {
	tokenizer := NewTokenizer(code)
	tokens := tokenizer.ScanTokens()
	for _, token := range tokens {
		token.String()
	}
}
