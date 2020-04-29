package engine

import (
	"regexp"
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


func shiftRune(runes []rune) (rune, []rune) {
	var r rune
	if len(runes) > 0 {
		r = runes[0]
		runes = runes[1:]
	}
	return r, runes
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}

// OPR symbols: - + - * / ( ) ^ % !
func isOperator(r string) bool {
	var validOp = regexp.MustCompile(`\(|\)|\-|\+|\/|\*|\^|\%`)
	return validOp.MatchString(r)
}
