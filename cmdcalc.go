package main

import (
	"bufio"
	"fmt"
	"os"
)

// processStack() : processes the string expression
func processStack(e string) (float64, error) {
	result := 0.0
	return result, nil
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("calc>")
		for scanner.Scan() {
			fmt.Println("exp:", scanner.Text())
			res, _ := processStack(scanner.Text())
			fmt.Println(res)
			fmt.Print("calc>")
		}
	}
}
