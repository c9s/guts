package coffeephp

import "testing"
import "coffeephp/codegen/phpcodegen"

func TestParser(t *testing.T) {
	inputs := []string{
		`pi = 3.1415926`,
		`pi = 3.1415926 + 4`,
		`pi = 4455
a = 4 * (2 + 3)
b = a`,
		`x = 3 * -2`,
		`if a > 10
	a = 10
`,
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
		visitor := phpcodegen.Visitor{}
		t.Log(visitor.Visit(parser.Val.val))
	}
}
