package main

import (
	"fmt"
	"gutscript"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage dump_lex [file]")
		os.Exit(0)
	}

	srcFile := os.Args[1]

	bytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lexer := gutscript.CreateLexer(string(bytes), 0)
	gutscript.DumpTokensFromChannel(lexer.Items)
}
