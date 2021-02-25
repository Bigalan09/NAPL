package main

import (
	"fmt"
	"os"

	"github.com/bigalan09/napl/scanner"
)

func main() {
	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}

	scanner := scanner.Scanner(file)
	for {
		pos, tok, lit := scanner.Scan()
		if tok == scanner.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	}
}
