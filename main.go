package main

// import "bufio"
// import "fmt"
// import "os"

func main() {
	// fi := bufio.NewReader(os.NewFile(0, "stdin"))
	lexer := CoffeeLex{
		input: "",
		items: make(chan LexItem),
	}
	CoffeeParse(&lexer)
}
