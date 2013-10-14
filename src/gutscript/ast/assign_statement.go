package ast

/*

*/
type AssignStatement struct {
	Variable Node
	Expr     Node
	Statement
}

func CreateAssignStatement(variable Node, expr Node) Node {
	return AssignStatement{
		Variable: variable,
		Expr:     expr,
	}
}
