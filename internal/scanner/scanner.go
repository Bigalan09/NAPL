package scanner

import (
	"bufio"
	"io"
)

// Token type
type Token int

// consts
const (
	// EOF type
	EOF = iota
	// ILLEGAL type
	ILLEGAL
	// IDENT type
	IDENT
	// INT type
	INT
	// SEMI type
	SEMI // ;
	// ADD type
	ADD // +
	// SUB type
	SUB // -
	// MUL type
	MUL // *
	// DIV type
	DIV // /
	// ASSIGN type
	ASSIGN // =
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",

	// Infix ops
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
}

func (t Token) String() string {
	return tokens[t]
}

// Position struct
type Position struct {
	line   int
	column int
}

// Lexer struct
type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

// Scanner reads in the file.
func Scanner(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Scan the file.
func (l *Lexer) Scan() (Position, Token, string) {
	for {
		_, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
		}

		panic(err)
	}
}
