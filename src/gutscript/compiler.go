package gutscript

import (
	"fmt"
	"gutscript/codegen/php"
	"io/ioutil"
)

func CompileFile(srcFile string) (string, error) {
	bytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return "", err
	}
	// code := string(bytes)
	lexer := GutsLex{
		input: string(bytes),
		start: 0,
		pos:   0,
		items: make(chan *GutsSymType, 100),
	}
	go lexer.run()
	parser := GutsParser{}
	if parser.Parse(&lexer) == 1 {
		return "", fmt.Errorf("syntax error")
	}
	lexer.close()
	visitor := php.Visitor{}
	// fmt.Printf("AST: %#v\n", parser.Val.val)
	return visitor.Visit(parser.Val.val), nil
}
