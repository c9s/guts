package gutscript

// vim:list:

import "testing"
import "gutscript/codegen/phpcodegen"
import "io/ioutil"
import "errors"
import "path/filepath"
import "fmt"

var parserInputs = []string{
	`pi = 3.1415926`,
	`pi = 3.1415926 + 4`,
	`pi = 4455
a = 4 * (2 + 3)
b = a`,

	// XXX: need trailing line
	`x = 3 * -2`,
	`if a > 10
    a = 10
`,
	`if a > 10
    a = 10
else
    a = 0
`,
	`if a > 10
    a = 10
    b = 20
else
    a = 0
    b = 0
`,
	// test if && single else if
	`if a > 10
    a = 10
    b = 20
elseif b > 10
    a = 1
    b = 1
`,
	// test if, elseif && else
	`if a > 10
    a = 10
    b = 20
elseif b > 10
    a = 1
    b = 1
else
    a = 0
    b = 0
`,
	`foo :: ->
    a = 10
    return a
`,
}

func LexFile(srcFile string) (chan *GutsSymType, error) {
	bytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}
	// code := string(bytes)
	lexer := GutsLex{
		input: string(bytes),
		start: 0,
		pos:   0,
		items: make(chan *GutsSymType, 100),
	}
	go lexer.run()
	return lexer.items, nil
}

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
		return "", errors.New("syntax error")
	}
	lexer.close()
	visitor := phpcodegen.Visitor{}
	fmt.Printf("AST: %#v\n", parser.Val.val)
	return visitor.Visit(parser.Val.val), nil
}

func TestCompileFile(t *testing.T) {
	srcFiles, err := filepath.Glob("tests/*.guts")
	if err != nil {
		t.Fatal(err)
	}
	for _, srcFile := range srcFiles {
		t.Log("Lexing", srcFile)
		items, err := LexFile(srcFile)
		if err != nil {
			t.Fatal(err)
		}
		if items != nil {
			dumpLexItems(items)
		}

		t.Log("Compiling", srcFile)
		out, err := CompileFile(srcFile)
		if err != nil {
			t.Fatalf("Compilation of %s failed: %s", srcFile, err)
		}
		t.Log(out)
		_ = out
	}
}

func TestParser(t *testing.T) {
	for _, input := range parserInputs {
		t.Log(input)
		lexer := GutsLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *GutsSymType, 100),
		}
		go lexer.run()
		parser := GutsParser{}
		if parser.Parse(&lexer) == 1 {
			t.Fatalf("syntax error: %s", input)
		}
		lexer.close()
	}
}
