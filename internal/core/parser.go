package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/exgene/forge/internal/tokenizer"
)

type Parser struct {
	tokens []tokenizer.Token
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) notEof() bool {
	if len(p.tokens) == 0 {
		return false
	}
	token := p.tokens[0]
	return token.Kind != tokenizer.EOF
}

func (p *Parser) at() tokenizer.Token {
	return p.tokens[0]
}

func (p *Parser) next() tokenizer.Token {
	token := p.tokens[0]
	p.tokens = p.tokens[1:]
	return token
}

func (p *Parser) expect(t tokenizer.TokenType, err_msg string) (tokenizer.Token, error) {
	token := p.tokens[0]
	p.tokens = p.tokens[1:]
	if token.Kind != t {
		return token, fmt.Errorf("Error::Parser::{%v}, Expected:{%v}====Got:{%v}\n", err_msg, t, token.Kind)
	}
	return token, nil
}

func (p *Parser) ProduceAST(source string) *Program {
	fmt.Print(source)
	t := tokenizer.NewTokenizer(source)
	p.tokens = t.ScanTokens()
	for _, token := range p.tokens {
		fmt.Println(token.String())
	}
	program := &Program{Body: []Statement{}}
	for p.notEof() {
		program.Body = append(program.Body, p.parseStatement())
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	return p.parseExpression()
}

func (p *Parser) parseExpression() Expression {
	return p.parseAdditiveExpression()
}

func (p *Parser) parseAdditiveExpression() Expression {
	left := p.parseMultiplicativeExpression()

	for p.notEof() && (p.at().Value == "+" || p.at().Value == "-") {
		operator := p.next().Value
		right := p.parseMultiplicativeExpression()
		left = &BinaryExpression{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}
	return left
}

func (p *Parser) parseMultiplicativeExpression() Expression {
	left := p.parsePrimaryExpression()

	for p.notEof() && (p.at().Value == "*" || p.at().Value == "/" || p.at().Value == "%") {
		operator := p.next().Value
		right := p.parsePrimaryExpression()
		left = &BinaryExpression{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}
	return left
}

func (p *Parser) parsePrimaryExpression() Expression {
	token := p.at()

	switch token.Kind {
	case tokenizer.Identifier:
		return &Identifier{Symbol: p.next().Value}
	case tokenizer.Numeric:
		val, err := strconv.ParseFloat(p.next().Value, 64)
		if err != nil {
			fmt.Printf("Error parsing the float value {%v}\n", token)
			os.Exit(1)
		}
		return &NumericLiteral{Value: val}
	case tokenizer.OpenParen:
		p.next()
		value := p.parseExpression()
		_, err := p.expect(tokenizer.ClosedParen, "Expected Closed parenthesis")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return value
	default:
		fmt.Printf("Idk what this token is bruh {%v}\n", token)
		os.Exit(1)
	}
	return nil
}
