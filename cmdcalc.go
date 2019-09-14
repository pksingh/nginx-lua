package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("calc>")
		for scanner.Scan() {
			fmt.Println("exp:", scanner.Text())
			fmt.Print("calc>")
		}
	}
}
