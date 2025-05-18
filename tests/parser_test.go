package tests

import (
	"reflect"
	"testing"

	"github.com/exgene/forge/internal/core"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *core.Program
	}{
		{
			name:  "Single number",
			input: "42",
			expected: &core.Program{
				Body: []core.Statement{
					&core.NumericLiteral{Value: 42.0},
				},
			},
		},
		{
			name:  "Simple addition",
			input: "5 + 10",
			expected: &core.Program{
				Body: []core.Statement{
					&core.BinaryExpression{
						Left:     &core.NumericLiteral{Value: 5.0},
						Right:    &core.NumericLiteral{Value: 10.0},
						Operator: "+",
					},
				},
			},
		},
		{
			name:  "Addition with precedence",
			input: "5 + 10 * 3",
			expected: &core.Program{
				Body: []core.Statement{
					&core.BinaryExpression{
						Left: &core.NumericLiteral{Value: 5.0},
						Right: &core.BinaryExpression{
							Left:     &core.NumericLiteral{Value: 10.0},
							Right:    &core.NumericLiteral{Value: 3.0},
							Operator: "*",
						},
						Operator: "+",
					},
				},
			},
		},
		{
			name:  "Parenthesized expression",
			input: "(5 + 10) * 3",
			expected: &core.Program{
				Body: []core.Statement{
					&core.BinaryExpression{
						Left: &core.BinaryExpression{
							Left:     &core.NumericLiteral{Value: 5.0},
							Right:    &core.NumericLiteral{Value: 10.0},
							Operator: "+",
						},
						Right:    &core.NumericLiteral{Value: 3.0},
						Operator: "*",
					},
				},
			},
		},
		{
			name:  "Identifier",
			input: "x",
			expected: &core.Program{
				Body: []core.Statement{
					&core.Identifier{Symbol: "x"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := core.NewParser()
			result := parser.ProduceAST(tt.input)

			if !compareAST(t, tt.expected, result) {
				t.Errorf("ASTs don't match for input: %s", tt.input)
			}
		})
	}
}

func compareAST(t *testing.T, expected, actual *core.Program) bool {
	if len(expected.Body) != len(actual.Body) {
		t.Errorf("Body length mismatch: expected %d, got %d", len(expected.Body), len(actual.Body))
		return false
	}

	for i, expectedStmt := range expected.Body {
		actualStmt := actual.Body[i]
		if !compareNode(t, expectedStmt, actualStmt) {
			t.Errorf("Statement %d mismatch", i)
			return false
		}
	}

	return true
}

func compareNode(t *testing.T, expected, actual core.Statement) bool {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Type mismatch: expected %T, got %T", expected, actual)
		return false
	}

	switch expectedNode := expected.(type) {
	case *core.NumericLiteral:
		actualNode, ok := actual.(*core.NumericLiteral)
		if !ok {
			return false
		}
		return expectedNode.Value == actualNode.Value

	case *core.Identifier:
		actualNode, ok := actual.(*core.Identifier)
		if !ok {
			return false
		}
		return expectedNode.Symbol == actualNode.Symbol

	case *core.BinaryExpression:
		actualNode, ok := actual.(*core.BinaryExpression)
		if !ok {
			return false
		}
		return expectedNode.Operator == actualNode.Operator &&
			compareNode(t, expectedNode.Left, actualNode.Left) &&
			compareNode(t, expectedNode.Right, actualNode.Right)

	case *core.Program:
		actualNode, ok := actual.(*core.Program)
		if !ok {
			return false
		}
		return compareAST(t, expectedNode, actualNode)

	default:
		t.Errorf("Unknown node type: %T", expected)
		return false
	}
}
