package coffeephp

import "testing"

var lextests = []struct {
	input    string
	expected TokenType
	cnt      int
}{
	{"a = 102", T_IDENTIFIER, 1},
	{"abc123 = 100", T_IDENTIFIER, 1},
	{"af = 102", T_NUMBER, 1},
	{"af = 102", T_ASSIGN, 1},
	{"af = 3.1415926", T_FLOATING, 1},
}

func expectLexInput(t *testing.T, input string, typ TokenType, cnt int) {
	lexer := CoffeeLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *CoffeeSymType, 100),
	}
	go lexer.run()

	var found int = 0
	var item *CoffeeSymType
	for {
		item = <-lexer.items
		if item == nil {
			break
		}
		t.Log(item)
		if item.typ == typ {
			found++
		}
		if item.typ == T_EOF {
			break
		}
	}
	if found != cnt {
		t.Fatalf("Expecting token %s x %d", typ, cnt)
	}
	lexer.close()
}

func TestLexerString(t *testing.T) {
	input := `a = "string content"`
	expectLexInput(t, input, T_STRING, 1)

	input = `a = "string content contains quote \""`
	expectLexInput(t, input, T_STRING, 1)
}

func TestLexerComment(t *testing.T) {
	input := `
	// oneline comment
	`
	expectLexInput(t, input, T_ONELINE_COMMENT, 1)

	input = `
	/* comment */
	`
	expectLexInput(t, input, T_COMMENT, 1)
}

func TestLexerAssignFloating(t *testing.T) {
	input := `
	b = 3.1415926
	bar = 478.123
	`
	expectLexInput(t, input, T_FLOATING, 2)
}

func BenchmarkLexer(b *testing.B) {
	input := `
	a = 102
	foo = 200
	pi = 3.1415926
	// oneline comment
	`
	var item *CoffeeSymType
	for i := 0; i < b.N; i++ {
		lexer := CoffeeLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *CoffeeSymType, 100),
		}
		go lexer.run()
		for {
			item = <-lexer.items
			if item == nil || item.typ == T_EOF {
				break
			}
		}
		lexer.close()
	}
}

func TestLexerAssign(t *testing.T) {
	for _, test := range lextests {
		expectLexInput(t, test.input, test.expected, test.cnt)
	}
}
