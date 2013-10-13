package ast

type IfStatement struct {
	Expr Expr
	Body *StatementNodeList
}

func CreateIfStatement(expr Expr, body *StatementNodeList) Node {
	return IfStatement{expr, body}
}
