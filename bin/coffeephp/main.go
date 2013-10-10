package main

import "coffeephp"
import _ "flag"

func main() {
	// XXX: Provide ParseFile function from coffeephp package.
	lexer := coffeephp.CoffeeLex{
		// XXX: take the input from file.
		input: "",
		start: 0,
		pos:   0,
		items: make(chan *coffeephp.CoffeeSymType, 100),
	}
	go lexer.run()
	coffeephp.CoffeeParse(lexer)
}
