package coffeephp

import "coffeephp"
import _ "flag"

func main() {
	lexer := coffeephp.CoffeeLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *CoffeeSymType, 100),
	}
	go lexer.run()
	coffeephp.CoffeeParse(lexer)
}
