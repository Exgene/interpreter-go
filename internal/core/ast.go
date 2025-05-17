package core

import "fmt"

type NodeType string

const (
	ProgramNode          NodeType = "program"
	NumericLiteralNode            = "numeric_literal"
	IdentifierNode                = "identifier"
	BinaryExpressionNode          = "binary_expression"
)

func PrintNode(node Statement, indent string) {
	switch n := node.(type) {
	case *Program:
		fmt.Println(indent + "Program")
		for _, stmt := range n.Body {
			PrintNode(stmt, indent+"  ")
		}
	case *BinaryExpression:
		fmt.Println(indent + "BinaryExpression (" + n.Operator + ")")
		fmt.Println(indent + "  Left:")
		PrintNode(n.Left, indent+"    ")
		fmt.Println(indent + "  Right:")
		PrintNode(n.Right, indent+"    ")
	case *Identifier:
		fmt.Println(indent + "Identifier: " + n.Symbol)
	case *NumericLiteral:
		fmt.Printf("%sNumericLiteral: %v\n", indent, n.Value)
	default:
		fmt.Printf("%sUnknown node: %T\n", indent, n)
	}
}

// Statement
type Statement interface {
	Kind() NodeType
}

type Program struct {
	Body []Statement
}

func (p *Program) Kind() NodeType { return ProgramNode }

// Expression
type Expression interface {
	Statement
	ExprNode()
}

// BinaryExpression
type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator string
}

func (b *BinaryExpression) Kind() NodeType {
	return BinaryExpressionNode
}
func (b *BinaryExpression) ExprNode() {}

// Identifier
type Identifier struct {
	Symbol string
}

func (i *Identifier) Kind() NodeType {
	return IdentifierNode
}
func (i *Identifier) ExprNode() {}

// NumericLiteral

type NumericLiteral struct {
	Value any
}

func (n *NumericLiteral) Kind() NodeType {
	return NumericLiteralNode
}
func (n *NumericLiteral) ExprNode() {}
