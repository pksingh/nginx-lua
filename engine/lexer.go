package engine

import (
	"fmt"
	"log"
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

// Tokenize performs a lexical analysis
func Tokenize(s string) []Token {
	runes := []rune(s)
	tokens := []Token{}

	for len(runes) > 0 {
		// Skip white space
		_, runes = readWhile(runes, isWhitespace)

		if len(runes) == 0 {
			break
		}

		r, shifted := shiftRune(runes)
		runes = shifted

		if isNumber(r) {
			s, readed := readWhile(runes, isNumber)
			runes = readed
			token := Token{
				Type:  NUM,
				Value: string(r) + s,
			}
			tokens = append(tokens, token)
		} else if isOperator(string(r)) {
			token := Token{
				Type:  OPR,
				Value: string(r),
			}
			tokens = append(tokens, token)
		} else {
			log.Fatalf("char %s not allowed", string(r))
		}
	}
	fmt.Println("tokens: ", tokens)
	return tokens
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
