package ast

import "strconv"

type Node interface {
}

type NumberNode struct {
	Val int64
}

type ExprNode struct {
}

type StatementNode struct {
}

type IfNode struct {
}

type ElseIfNode struct {
}

type ElseNode struct {
}

func CreateExprNode() {

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
