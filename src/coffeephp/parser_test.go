package coffeephp

import "testing"
import "coffeephp/codegen/phpcodegen"

func TestParser(t *testing.T) {
	inputs := []string{
		`pi = 3.1415926`,
		`pi = 3.1415926 + 4`,
		`pi = 4455
a = 1 + (2 * 3)
b = a`,
	}
	for _, input := range inputs {
		t.Log(input)
		lexer := CoffeeLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *CoffeeSymType, 100),
		}
		go lexer.run()

		parser := CoffeeParser{}
		if parser.Parse(&lexer) == 1 {
			t.Fatal("syntax error")
		}
		lexer.close()
		t.Logf("AST: %#v", parser.Val.val)

		t.Log(phpcodegen.Visit(parser.Val.val))
	}
}
