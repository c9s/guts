package gutscript

// vim:list:

import "testing"
import "io/ioutil"
import "path/filepath"

var parserInputs = []string{
	`pi = 3.1415926`,
	`pi = 3.1415926 + 4`,
	`pi = 4455
a = 4 * (2 + 3)
b = a`,

	// XXX: need trailing line
	`x = 3 * -2`,
}

func LexFile(srcFile string) (chan *GutsSymType, error) {
	bytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}

	lexer := CreateLexer(string(bytes), 0)
	return lexer.Items, nil
}

func TestCompileFile(t *testing.T) {
	srcFiles, err := filepath.Glob("tests/*.guts")
	if err != nil {
		t.Fatal(err)
	}
	for _, srcFile := range srcFiles {
		t.Log("Compiling", srcFile)
		out, err := CompileFile(srcFile)
		if err != nil {
			t.Fatalf("Compilation of %s failed: %s", srcFile, err)
		}
		t.Log("Compilation success.", srcFile)
		t.Log("Compilation result:")
		t.Log(out)
	}
}

func TestParser(t *testing.T) {
	for i, input := range parserInputs {
		t.Logf("Testing test case %d: %s", i, input)
		lexer := CreateLexer(input, 100)
		parser := GutsParser{}
		if parser.Parse(lexer) == 1 {
			t.Fatalf("syntax error: %s", input)
		}
		lexer.close()
	}
}
