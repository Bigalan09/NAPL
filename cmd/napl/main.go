package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}

	scanner := Scanner(file)
	for {
		pos, tok, lit := scanner.Scan()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	}
}
