package tests

import (
	"testing"

	"github.com/exgene/forge/internal/tokenizer"
)

func TestTokenizer_SimpleExpressions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []tokenizer.Token
	}{
		{
			name:  "Single number",
			input: "42",
			expected: []tokenizer.Token{
				{Kind: tokenizer.Numeric, Value: "42"},
				{Kind: tokenizer.EOF, Value: ""},
			},
		},
		{
			name:  "Addition",
			input: "42 + 5",
			expected: []tokenizer.Token{
				{Kind: tokenizer.Numeric, Value: "42"},
				{Kind: tokenizer.BinaryOperator, Value: "+"},
				{Kind: tokenizer.Numeric, Value: "5"},
				{Kind: tokenizer.EOF, Value: ""},
			},
		},
		{
			name:  "Multiplication",
			input: "42 * 5",
			expected: []tokenizer.Token{
				{Kind: tokenizer.Numeric, Value: "42"},
				{Kind: tokenizer.BinaryOperator, Value: "*"},
				{Kind: tokenizer.Numeric, Value: "5"},
				{Kind: tokenizer.EOF, Value: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tkn := tokenizer.NewTokenizer(tt.input)
			tokens := tkn.ScanTokens()

			if len(tokens) != len(tt.expected) {
				t.Errorf("Expected %d tokens, got %d", len(tt.expected), len(tokens))
				t.Errorf("Expected: %v", tt.expected)
				t.Errorf("Got: %v", tokens)
				return
			}

			for i, token := range tokens {
				if token.Kind != tt.expected[i].Kind || token.Value != tt.expected[i].Value {
					t.Errorf("Token %d mismatch:\nExpected: %v\nGot: %v", i, tt.expected[i], token)
				}
			}
		})
	}
}
