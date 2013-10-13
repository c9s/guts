package ast

import "strconv"

type FloatingNumber struct {
	Val float64
}

func CreateFloatingNumber(token string) Node {
	var err error
	val, err := strconv.ParseFloat(token, 64)
	if err != nil {
		panic(err)
	}
	n := FloatingNumber{}
	n.Val = val
	return n
}
