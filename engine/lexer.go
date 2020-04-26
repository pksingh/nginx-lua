package engine

import (
	"unicode"
)

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

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}
