package ast

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
