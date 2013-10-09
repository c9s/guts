package main

import "testing"

func TestLexer(t *testing.T) {
	input := `
	a = 1
	b = 2
	foo23 = 3
	`
	lexer := CoffeeLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan LexItem),
	}
	for r := lexer.next(); r != eof; r = lexer.next() {
		t.Log(r)
		_ = r
	}
	_ = lexer
}
