package engine

type isFunc func(rune) bool

// LexerType lexer type
type LexerType int

const (
	OPR LexerType = iota // Operator Symbol
	NUM                  // Number/Digit Character
)

// Token define a token
type Token struct {
	Type  LexerType
	Value string
}
