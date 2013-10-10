package ast

type StatementNode struct {
}

func CreateStatementNode() Node {
	return StatementNode{}
}

func CreateExprStatementNode(node Node) Node {
	return StatementNode{}
}

func CreateStatementNodeList() []StatementNode {
	return []StatementNode{}
}
