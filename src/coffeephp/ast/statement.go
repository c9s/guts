package ast

type StatementNode struct {
}

func CreateStatement() Node {
	return StatementNode{}
}

func CreateExprStatement(node Node) Node {
	return StatementNode{}
}
