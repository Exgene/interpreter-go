package core

type NodeType string

const (
	ProgramNode          NodeType = "program"
	NumericLiteralNode            = "numeric_literal"
	IdentifierNode                = "identifier"
	BinaryExpressionNode          = "binary_expression"
)

// Statement
type Statement interface {
	Kind() NodeType
}

type Program struct {
	body []Statement
}
func (p *Program) Kind() NodeType { return ProgramNode }

// Expression
type Expression interface {
	Statement
	ExprNode()
}

// BinaryExpression
type BinaryExpression struct {
	left     Expression
	right    Expression
	operator string
}

func (b *BinaryExpression) Kind() NodeType {
	return BinaryExpressionNode
}
func (b *BinaryExpression) ExprNode()

// Identifier
type Identifier struct {
	symbol string
}

func (i *Identifier) Kind() NodeType {
	return IdentifierNode
}
func (i *Identifier) ExprNode()

// NumericLiteral

type NumericLiteral struct {
	value any
}

func (n *NumericLiteral) Kind() NodeType {
	return NumericLiteralNode
}
func (n *NumericLiteral) ExprNode()
