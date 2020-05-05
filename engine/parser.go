package engine

// PrecedenceOperators define Precedence of operators
var PrecedenceOperators = map[string]int{
	"(": 10,
	"+": 20,
	"-": 20,
	"/": 30,
	"*": 30,
	"^": 40,
	"%": 50,
}

// ShiftToken shift token array
func ShiftToken(tokens []Token) (Token, []Token) {
	var t Token
	if len(tokens) > 0 {
		t = tokens[0]
		tokens = tokens[1:]
	}
	return t, tokens
}

