package tokenizer

import "fmt"

type TokenType int

const (
	Numeric    TokenType = iota
	Identifier TokenType = iota
)

type Token struct {
	kind  TokenType
	value *float32
}

func TokenizeShit(cd string) {
	t := Token{kind: TokenType(Numeric), value: nil}
	fmt.Printf("Token:{%d}=={}", t.kind)
}
