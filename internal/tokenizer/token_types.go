package tokenizer

type TokenType int

const (
	Numeric TokenType = iota
	Identifier
	Equals
	OpenParen
	ClosedParen
	OpenCurlyBracket
	ClosedCurlyBracket
	BinaryOperator
	Let
	Const
	EOF
	String
)

type Token struct {
	Kind  TokenType
	Value any
}

type Tokenizer struct {
	idx      int
	buf      string
	tokens   []Token
	src_code string
}
