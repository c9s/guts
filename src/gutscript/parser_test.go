package gutscript

// vim:list:

import "testing"
import "gutscript/codegen/phpcodegen"
import "io/ioutil"
import "errors"
import "path/filepath"

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

func LexFile(srcFile string) error {
	bytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}
	// code := string(bytes)
	lexer := GutsLex{
		input: string(bytes),
		start: 0,
		pos:   0,
		items: make(chan *GutsSymType, 100),
	}
	go lexer.run()
	dumpLexItems(lexer.items)
	return nil
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
	return visitor.Visit(parser.Val.val), nil
}

func TestCompileFile(t *testing.T) {
	srcFiles, err := filepath.Glob("tests/*.guts")
	if err != nil {
		t.Fatal(err)
	}
	for _, srcFile := range srcFiles {
		t.Log("Lexing", srcFile)
		LexFile(srcFile)

		t.Log("Compiling", srcFile)
		out, err := CompileFile(srcFile)
		if err != nil {
			t.Fatalf("Compilation of %s failed: %s", srcFile, err)
		}
		t.Log(out)
		_ = out
	}
}

func TestCodeGen(t *testing.T) {
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
			t.Fatal("syntax error")
		}
		lexer.close()
		t.Logf("AST: %#v", parser.Val.val)
		visitor := phpcodegen.Visitor{}
		t.Log(visitor.Visit(parser.Val.val))
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
			t.Fatal("syntax error")
		}
		lexer.close()
	}
}
