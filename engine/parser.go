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
