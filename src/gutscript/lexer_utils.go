package gutscript

import "fmt"

func CreateLexer(input string, capacity uint32) *GutsLex {
	if capacity == 0 {
		// default capacity
		capacity = 100
	}
	lexer := GutsLex{
		Input: input,
		Start: 0,
		Pos:   0,
		Items: make(chan *GutsSymType, capacity),
	}
	go lexer.run()
	return &lexer
}

func DumpTokenTypeList(items []*GutsSymType) string {
	var out string = "[]TokenType{\n"
	for _, item := range items {
		out += "\t" + GetTokenName(int(item.typ)) + ",\n"
	}
	out += "}\n"
	return out
}

// Given a token type list, check the returned token items
func DumpTokensFromChannel(itemChannel chan *GutsSymType) {
	var i = 0
	var item *GutsSymType
	for {
		item = <-itemChannel
		if item == nil {
			break
		}
		if item.typ == T_EOF || item.typ == eof {
			break
		}
		fmt.Printf("%s: '%s'\n", GetTokenName(int(item.typ)), item.val)
		// fmt.Printf("'%s' : %s\n", item.val, GetTokenName(int(item.typ)))
		i++
	}
	fmt.Println("==========================")
	fmt.Printf("%d tokens\n", i)
}
