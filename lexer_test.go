package main

import "testing"

func testLexInput(t *testing.T, input string) {
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
}

func TestLexerComment(t *testing.T) {
	input := `
	// oneline comment
	`
	testLexInput(t, input)

	input = `
	/* comment */
	`
	testLexInput(t, input)
}

func TestLexerAssign(t *testing.T) {
	input := `
	a = 102
	b = 3.1415926
	foo = 200
	bar = 478.123
	`
	testLexInput(t, input)
}
