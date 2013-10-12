package ast

type StatementNode struct {
}

func CreateStatementNode() Node {
	return StatementNode{}
}

func CreateExprStatementNode(node Node) Node {
	return StatementNode{}
}

type StatementNodeList []Node

func (self *StatementNodeList) Append(node Node) {
	old := *self
	newList := append(old, node)
	*self = newList
}

func CreateStatementNodeList() *StatementNodeList {
	return &StatementNodeList{}
}
