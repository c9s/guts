package ast

type StatementNode struct {
}

func CreateStatementNode() Node {
	return StatementNode{}
}

func CreateExprStatementNode(node Node) Node {
	return StatementNode{}
}

type NodeList []Node

func (self *NodeList) Append(node Node) {
	old := *self
	newList := append(old, node)
	*self = newList
}

func CreateStatementNodeList() *NodeList {
	return &NodeList{}
}
