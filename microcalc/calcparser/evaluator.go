package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	. "engine" // will refer to github.com/pksingh/uservices/engine
)

// Evaluate postfix expression
func Evaluate(tokens []Token) ResultResponse {
	result := ResultResponse{}

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

			res, err := DoPost(t.Value, r, l)
			if err != nil {
				log.Panicln("Error: ", err.Error())
			}
			nr := NodeResponse{}
			if err := json.NewDecoder(res.Body).Decode(&nr); err != nil {
				fmt.Println("error : ", err.Error())
			}
			fmt.Println("res : ", res)
			fmt.Println("nr res : ", nr)
			calc := nr.Result
			result.Origins = append(result.Origins, nr)
			stack = append(stack, calc)
			fmt.Printf("Evaluate: %d %s %d = %d\n", r, t.Value, l, calc)
		}
	}

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
