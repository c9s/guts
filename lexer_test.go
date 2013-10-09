package main

import "testing"

func TestLexer(t *testing.T) {
	input := `
	a = 102
	b = 3.1415926
	`
	lexer := CoffeeLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *CoffeeSymType, 100),
	}
	go lexer.run()

	var item *CoffeeSymType
	for {
		item = <-lexer.items
		t.Log(item)
		if item.typ == T_EOF {
			break
		}
	}
	lexer.close()

	/*
		var lval CoffeeSymType
		for lexer.Lex(&lval) != eof {
			t.Log(lval)
		}

		for r := lexer.next(); r != eof; r = lexer.next() {
			t.Log(r)
			_ = r
		}
	*/
}
