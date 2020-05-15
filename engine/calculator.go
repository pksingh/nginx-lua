package engine

import (
	"fmt"
	"log"
	"strconv"
)

// Calculate postfix expression
func Calculate(tokens []Token) int {
	result := 0

	stack := []int{}

	for len(tokens) > 0 {
		t, shifted := ShiftToken(tokens)
		tokens = shifted

		if t.Type == NUM {
			intValue, err := strconv.Atoi(t.Value)
			if err != nil {
				log.Fatalf("cast error %s", t.Value)
			}
			stack = append(stack, intValue)
		} else {
			// Otherwise we got OPR, get operands
			l, lPoped := popInt(stack)
			stack = lPoped

			r, rPoped := popInt(stack)
			stack = rPoped

			var calc int
			switch t.Value {
			case "+":
				calc = r + l
				stack = append(stack, calc)
			case "-":
				calc = r - l
				stack = append(stack, calc)
			case "*":
				calc = r * l
				stack = append(stack, calc)
			case "/":
				calc = r / l
				stack = append(stack, calc)
			case "%":
				calc = r % l
				stack = append(stack, calc)
			}
			fmt.Printf("Evaluate: %d %s %d = %d\n", r, t.Value, l, calc)
		}
	}

	result, _ = popInt(stack)

	return result
}

func popInt(values []int) (int, []int) {
	var v int
	if len(values) > 0 {
		v = values[len(values)-1]
		values = values[:len(values)-1]
	}
	return v, values
}
