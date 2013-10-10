package coffeephp

import "testing"

func TestParser(t *testing.T) {
	inputs := []string{
		`pi = 3.1415926`,
		`pi = 3.1415926 + 4`,
	}
	for _, input := range inputs {
		lexer := CoffeeLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *CoffeeSymType, 100),
		}
		go lexer.run()
		ret := CoffeeParse(&lexer)
		lexer.close()
		if ret == 1 {
			t.Fatal("syntax error")
		}
	}
}
