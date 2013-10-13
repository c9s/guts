package ast

type UnaryExpr struct {
	Op  byte
	Val Node
}

type Expr struct {
	Op          byte
	Left        Node
	Right       Node
	Parenthesis bool
}

func CreateExpr(op byte, left, right Node) Node {
	return Expr{
		Op:    op,
		Left:  left,
		Right: right,
	}
}
