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
	token := p.tokens[0]
// fmt.Printf("Im called with token %v\n", token)
	// fmt.Printf("I resolve to ==> %v \n", token.Kind != tokenizer.EOF)
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
	t := tokenizer.NewTokenizer(source)
	p.tokens = t.ScanTokens()
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

	for p.at().Value == "+" || p.at().Value == "-" {
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

	for p.at().Value == "*" || p.at().Value == "/" || p.at().Value == "%" {
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
	token := p.next()

	switch token.Kind {
	case tokenizer.Identifier:
		return &Identifier{Symbol: token.Value}
	case tokenizer.Numeric:
		val, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			fmt.Printf("Error parsing the float value {%v}\n", token)
			os.Exit(1)
		}
		return &NumericLiteral{Value: val}
	case tokenizer.OpenParen:
		value := p.parseExpression()
		p.expect(tokenizer.ClosedParen, "Expected Closed parenthesis")
		return value
	case tokenizer.EOF:
		return nil	
	default:
		fmt.Printf("Idk what this token is bruh {%v}\n", token)
		os.Exit(1)
	}
	return nil
}
