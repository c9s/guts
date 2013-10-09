package main

import "bufio"
import "fmt"
import "os"

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			lexer := CoffeeLexToken{
				input: eqn,
				items: make(chan LexItem),
			}
			CoffeeParse(&lexer)
		} else {
			break
		}
	}
}
