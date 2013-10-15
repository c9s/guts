package ast

type IfElseStatement struct {
	Expr     Expr
	Body     *StatementList
	ElseBody *StatementList
}

func CreateIfElseStatement(expr Expr, body *StatementList, elseBody *StatementList) Node {
	return IfElseStatement{expr, body, elseBody}
}
