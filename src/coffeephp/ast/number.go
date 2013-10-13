package ast

import "strconv"

type Number struct {
	Val int64
}

func CreateNumber(token string) Node {
	var err error
	var base int = 10
	if token[0] == '0' {
		base = 8
	}
	val, err := strconv.ParseInt(token, base, 64)
	if err != nil {
		panic(err)
	}
	n := Number{}
	n.Val = val
	return n
}
