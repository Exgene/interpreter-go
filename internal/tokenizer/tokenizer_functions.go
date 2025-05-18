package tokenizer

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func (tokenizer *Tokenizer) peek(ahead int) rune {
	if tokenizer.idx+ahead > len(tokenizer.src_code) {
		// IM USING THIS AS AN ASSERT UNTIL I FIGURE OUT GO!!!!
		fmt.Println("Shouldnt happen")
	}
	c, _ := utf8.DecodeRuneInString(tokenizer.src_code[tokenizer.idx+ahead-1:])
	return c
}

func (t *Tokenizer) isAtEnd() bool {
	if t.idx >= len(t.src_code) {
		return true
	}
	return false
}

func (t *Tokenizer) scanToken() error {
	for !t.isAtEnd() && is_skippable(t.peek(1)) {
		t.next()
	}

	if t.isAtEnd() {
		return nil
	}

	c := t.next()

	switch {
	case c == '(':
		t.addToTokensArray(Token{Kind: TokenType(OpenParen), Value: string(c)})
	case c == ')':
		t.addToTokensArray(Token{Kind: TokenType(ClosedParen), Value: string(c)})
	case c == '+' || c == '-' || c == '/' || c == '*' || c == '%':
		t.addToTokensArray(Token{Kind: TokenType(BinaryOperator), Value: string(c)})
	case c == '=':
		t.addToTokensArray(Token{Kind: TokenType(Equals), Value: string(c)})
	case c == '{':
		t.addToTokensArray(Token{Kind: TokenType(OpenCurlyBracket), Value: string(c)})
	case c == '}':
		t.addToTokensArray(Token{Kind: TokenType(ClosedCurlyBracket), Value: string(c)})
	case c == '"':
		t.buf = ""
		for !t.isAtEnd() && t.peek(1) != '"' {
			t.buf += string(t.next())
		}

		if t.isAtEnd() {
			return fmt.Errorf("Missing \" value in the string, buf: {%v}", t.buf)
		}
		t.addToTokensArray(Token{Kind: TokenType(String), Value: t.buf})
		t.buf = ""
	case unicode.IsNumber(c):
		t.scanNumber(c)
	case unicode.IsLetter(c):
		t.scanLiteral(c)
	}

	return nil
}

func (t *Tokenizer) addToTokensArray(token Token) {
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) scanLiteral(c rune) {
	t.buf += string(c)
	for !t.isAtEnd() && unicode.IsLetter(t.peek(1)) {
		t.buf += string(t.next())
	}

	tokenType, isBufAKeyword := isKeyword(t.buf)

	if isBufAKeyword {
		t.addToTokensArray(Token{Kind: tokenType, Value: t.buf})
	} else {
		t.addToTokensArray(Token{Kind: TokenType(Identifier), Value: t.buf})
	}

	t.buf = ""
}

func (t *Tokenizer) scanNumber(c rune) {
	t.buf += string(c)
	for !t.isAtEnd() && unicode.IsNumber(t.peek(1)) {
		t.buf += string(t.next())
	}
	t.addToTokensArray(Token{Kind: Numeric, Value: t.buf})
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
		err := t.scanToken()
		if err != nil {
			fmt.Printf("Error while parsing tokens...%v", err)
		}
	}

	t.addToTokensArray(Token{TokenType(EOF), ""})
	return t.tokens
}
