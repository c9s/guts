package ast

import "strconv"

type NumberNode struct {
	Val int64
}

func CreateNumberNode(token string) Node {
	var err error
	var base int = 10
	if token[0] == '0' {
		base = 8
	}
	val, err := strconv.ParseInt(token, base, 64)
	if err != nil {
		panic(err)
	}
	n := NumberNode{}
	n.Val = val
	return n
}
