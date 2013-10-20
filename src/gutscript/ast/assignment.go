package ast

/*

*/
type Assignment struct {
	Variable Node
	Expr     Node
	Statement
}

func CreateAssignment(variable Node, expr Node) Node {
	return Assignment{
		Variable: variable,
		Expr:     expr,
	}
}
