package ast

type IfElseStatement struct {
	Expr     Expr
	Body     *StatementNodeList
	ElseBody *StatementNodeList
}

func CreateIfElseStatement(expr Expr, body *StatementNodeList, elseBody *StatementNodeList) Node {
	return IfElseStatement{expr, body, elseBody}
}
