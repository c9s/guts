package gutscript

import "fmt"

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
		fmt.Printf("token %s: '%s'\n", GetTokenName(int(item.typ)), item.val)
		i++
	}
	fmt.Printf("%d tokens\n", i)
}
