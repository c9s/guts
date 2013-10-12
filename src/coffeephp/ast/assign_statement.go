package ast

/*

*/
type AssignStatementNode struct {
	Variable Node
	Expr     Node
	StatementNode
}

func CreateAssignStatementNode(variable Node, expr Node) Node {
	return AssignStatementNode{
		Variable: variable,
		Expr:     expr,
	}
}
