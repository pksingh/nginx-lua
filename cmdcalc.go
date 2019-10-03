package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseArgs() : validate Arguments by parsing as number
func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

// processStack() : processes the string expression
func processStack(e string) (float64, error) {
	result := 0.0
	fmt.Printf("Passed Arg: %#v, Len: %d\n", e, len(e))

	c := strings.Split(e, " ")
	fmt.Println("Arg:", e, " len:", len(c))
	if len(c)-1 < 2 {
		return 0.0, errors.New("error: missing arguments (pass space between number and symbols)")
	}
	num1, num2, err := parseArgs(c)
	if err != nil {
		return 0.0, err
	}
	fmt.Println("num1:", num1, " num2:", num2)

	switch c[1] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero.")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result, nil
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("calc>")
		for scanner.Scan() {
			fmt.Println("exp:", scanner.Text())
			res, err := processStack(scanner.Text())
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			fmt.Print("calc>")
		}
	}
}
