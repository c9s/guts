package main

import "gutscript"
import _ "flag"

func main() {
	// XXX: Provide ParseFile function from gutscript package.
	lexer := gutscript.GutsLex{
		// XXX: take the input from file.
		input: "",
		start: 0,
		pos:   0,
		items: make(chan *gutscript.GutsSymType, 100),
	}
	go lexer.run()
	gutscript.GutsParse(lexer)
}
