package core

import (
	"github.com/exgene/forge/internal/tokenizer"
)

type Parser struct {
	tokens []tokenizer.Token
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) notEof() bool {
	token := p.tokens[0]
	return token.Kind != tokenizer.TokenType(tokenizer.EOF)
}
