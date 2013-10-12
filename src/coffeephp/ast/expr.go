package ast

type UnaryExprNode struct {
	Op  byte
	Val Node
}

type ExprNode struct {
	Op          byte
	Left        Node
	Right       Node
	Parenthesis bool
}

func CreateExprNode(op byte, left, right Node) Node {
	return ExprNode{
		Op:    op,
		Left:  left,
		Right: right,
	}
}
