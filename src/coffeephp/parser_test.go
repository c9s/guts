package coffeephp

import "testing"

func TestParser(t *testing.T) {
	input := "pi = 3.1415926"
	lexer := CoffeeLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *CoffeeSymType, 100),
	}
	go lexer.run()
	CoffeeParse(&lexer)
	lexer.close()
}
