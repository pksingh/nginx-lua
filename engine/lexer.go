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


func readWhile(runes []rune, predicate isFunc) (string, []rune) {
	var s = ""
	for len(runes) > 0 && predicate(runes[0]) {
		c, shifted := shiftRune(runes)
		runes = shifted
		s += string(c)
	}
	return s, runes
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
