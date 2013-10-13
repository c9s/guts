package ast

type StatementNodeList []Node

func (self *StatementNodeList) Append(node Node) {
	old := *self
	newList := append(old, node)
	*self = newList
}

func CreateStatementList() *StatementNodeList {
	return &StatementNodeList{}
}
