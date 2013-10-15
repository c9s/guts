package ast

type IfStatement struct {
	Expr       Expr
	Body       *StatementList
	ElseIfList []ElseIfStatement
	ElseBody   *StatementList
}

type ElseIfStatement struct {
	Expr Expr
	Body *StatementList
}

func CreateIfStatement(expr Expr, body *StatementList) Node {
	return &IfStatement{expr, body, []ElseIfStatement{}, nil}
}

func (self *IfStatement) AddElseIf(expr Expr, body *StatementList) {
	self.ElseIfList = append(self.ElseIfList, ElseIfStatement{expr, body})
}

func (self *IfStatement) SetElse(body *StatementList) {
	self.ElseBody = body
}
