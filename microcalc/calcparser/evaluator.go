package main

import (
	. "engine" // will refer to github.com/pksingh/uservices/engine
)

// Evaluate postfix expression
func Evaluate(tokens []Token) ResultResponse {
	result := ResultResponse{}

	stack := []int{}

	result.Result, _ = popInt(stack)

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
