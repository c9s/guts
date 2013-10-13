package ast

type IfStatement struct {
	Expr       Expr
	Body       *StatementNodeList
	ElseIfList []ElseIfStatement
	ElseBody   *StatementNodeList
}

type ElseIfStatement struct {
	Expr Expr
	Body *StatementNodeList
}

func CreateIfStatement(expr Expr, body *StatementNodeList) Node {
	return &IfStatement{expr, body, []ElseIfStatement{}, nil}
}

func (self *IfStatement) AddElseIf(expr Expr, body *StatementNodeList) {
	self.ElseIfList = append(self.ElseIfList, ElseIfStatement{expr, body})
}

func (self *IfStatement) SetElse(body *StatementNodeList) {
	self.ElseBody = body
}
