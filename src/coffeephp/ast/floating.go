package ast

import "strconv"

type FloatingNumberNode struct {
	Val float64
}

func CreateFloatingNumberNode(token string) Node {
	var err error
	val, err := strconv.ParseFloat(token, 64)
	if err != nil {
		panic(err)
	}
	n := FloatingNumberNode{}
	n.Val = val
	return n
}
