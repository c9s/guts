package gutscript

// vim:list:

import "testing"
import "fmt"

var lextests = []struct {
	input    string
	expected TokenType
	cnt      int
}{
	{"a = 102", T_IDENTIFIER, 1},
	{"a = 102\n", T_NEWLINE, 0},
	{"a = 102\nb = 333\n", T_NEWLINE, 1},
	{"abc123 = 100", T_IDENTIFIER, 1},
	{"af = 102", T_NUMBER, 1},
	{"af = 102", '=', 1},
	{"af = 3.1415926", T_FLOATING, 1},
	{"af = bk", T_IDENTIFIER, 2},
	{"b = 3 * -10", T_IDENTIFIER, 1},
	{"// oneline comment", T_ONELINE_COMMENT, 1},
	{"/* comment */", T_COMMENT, 1},
	{`a = "string content"`, T_STRING, 1},
	{`a = "string content contains quote \""`, T_STRING, 1},
	{`if a > 3
    a = 10
    b = 20
`, T_IF, 1},
	{`if a > 3
    a = 10
    b = 20
else
    a = 1
`, T_ELSE, 1},
	{`if a > 3
    a = 10
    b = 20
elseif a < 10
    b = 2
else
    a = 1
`, T_ELSEIF, 1},
	{`foo :: ->
    a = 10
    b = 20
`, T_FUNCTION_PROTOTYPE, 1},
}

// Given a token type list, check the returned token items
func expectLexItems(t *testing.T, itemChannel chan *GutsSymType, expectedItems []TokenType) {
	var i = 0
	var item *GutsSymType

	if len(expectedItems) == 0 {
		var items []*GutsSymType
		for {
			item = <-itemChannel
			if item == nil {
				break
			}
			if item.typ == T_EOF || item.typ == eof {
				break
			}
			items = append(items, item)
		}
		fmt.Println(DumpTokenTypeList(items))
		return
	}

	for {
		item = <-itemChannel
		if item == nil {
			break
		}
		if item.typ == T_EOF || item.typ == eof {
			break
		}
		if i+1 > len(expectedItems) {
			t.Log("Given more tokens than expected tokens.")
			return
			// break
		}

		if item.typ != expectedItems[i] {
			t.Fatalf("Expecting token %s but we got %s",
				GetTokenName(int(expectedItems[i])),
				GetTokenName(int(item.typ)))
		}
		t.Logf("ok: token %s: '%s'\n", GetTokenName(int(item.typ)), item.val)
		i++
	}
	if i < len(expectedItems) || i > len(expectedItems) {
		t.Fatalf("Expecting %d tokens, but we got %d tokens", len(expectedItems), i)
	}
	t.Logf("ok: %d tokens matched.\n", len(expectedItems))
}

func expectLexInput(t *testing.T, input string, typ TokenType, cnt int) {
	lexer := GutsLex{
		input: input,
		start: 0,
		pos:   0,
		items: make(chan *GutsSymType, 100),
	}
	go lexer.run()

	var found int = 0
	var item *GutsSymType
	for {
		item = <-lexer.items
		if item == nil {
			break
		}
		t.Logf("Got token %s: %s", GetTokenName(int(item.typ)), item.val)
		if item.typ == typ {
			found++
		}
		if item.typ == T_EOF || item.typ == eof {
			break
		}
	}
	if found != cnt {
		t.Fatalf("Expecting token %s x %d", GetTokenName(int(typ)), cnt)
	}
	lexer.close()
}

func BenchmarkLexer(b *testing.B) {
	input := `
	a = 102
	foo = 200
	pi = 3.1415926
	// oneline comment
	`
	var item *GutsSymType
	for i := 0; i < b.N; i++ {
		lexer := GutsLex{
			input: input,
			start: 0,
			pos:   0,
			items: make(chan *GutsSymType, 100),
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

type LexTestingItem struct {
	File   string
	Tokens []TokenType
}

var lexInputFiles []LexTestingItem = []LexTestingItem{
	LexTestingItem{"tests/01_assignment.guts", []TokenType{T_IDENTIFIER, '=', T_FLOATING}},
	LexTestingItem{"tests/02_assignment_expr.guts", []TokenType{
		T_IDENTIFIER, '=', T_NUMBER, T_NEWLINE,
		T_IDENTIFIER, '=', T_IDENTIFIER, '*', T_NUMBER, '+', T_NUMBER,
	}},
	LexTestingItem{"tests/05_function_call_01.guts", []TokenType{
		T_IDENTIFIER,
		T_FUNCTION_PROTOTYPE,
		'(',
		T_IDENTIFIER,
		')',
		T_FUNCTION_GLYPH,
		T_INDENT,
		T_RETURN,
		T_IDENTIFIER,
		'*',
		T_IDENTIFIER,
		T_OUTDENT,
		T_NEWLINE,
		T_IDENTIFIER,
		'(',
		T_FLOATING,
		')',
	}},
	LexTestingItem{"tests/05_function_call_02.guts", []TokenType{
		T_IDENTIFIER,
		T_FUNCTION_PROTOTYPE,
		'(',
		T_IDENTIFIER,
		')',
		T_FUNCTION_GLYPH,
		T_INDENT,
		T_RETURN,
		T_IDENTIFIER,
		'*',
		T_IDENTIFIER,
		T_OUTDENT,
		T_NEWLINE,
		T_IDENTIFIER,
		'(',
		T_NUMBER,
		')',
	}},
}

func TestLexerShortSample(t *testing.T) {
	for _, test := range lextests {
		t.Log("Testing lex sample:", test.input)
		expectLexInput(t, test.input, test.expected, test.cnt)
	}
}

func TestLexerSample(t *testing.T) {
	for _, testItem := range lexInputFiles {
		t.Log("Testing", testItem.File)
		items, err := LexFile(testItem.File)
		if err != nil {
			t.Fatal(err)
		}
		if items != nil {
			expectLexItems(t, items, testItem.Tokens)
		}
	}
}
