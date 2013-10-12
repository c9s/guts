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

/*

*/
type AssignStatementNode struct {
	Assign Node
	Expr   Node
}

func CreateAssignStatementNode(assign Node, expr Node) Node {
	return AssignStatementNode{
		Assign: assign,
		Expr:   expr,
	}
}
