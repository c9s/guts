package ast

type UnaryExpr struct {
	Op  int
	Val Node
}

type Expr struct {
	Op          int
	Left        Node
	Right       Node
	Parenthesis bool
}

func CreateExpr(op int, left, right Node) Node {
	return Expr{
		Op:    op,
		Left:  left,
		Right: right,
	}
}
