package ast

type StatementList []Node

func (self *StatementList) Append(node Node) {
	old := *self
	newList := append(old, node)
	*self = newList
}

func CreateStatementList() *StatementList {
	return &StatementList{}
}
