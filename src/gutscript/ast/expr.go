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

func CreateUnaryExpr(op int, val Node) UnaryExpr {
	return UnaryExpr{op, val}
}

func CreateExpr(op int, left, right Node) Expr {
	return Expr{
		Op:    op,
		Left:  left,
		Right: right,
	}
}
