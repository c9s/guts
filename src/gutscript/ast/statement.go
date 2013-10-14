package ast

type Statement struct {
}

type ExprStatement struct {
	Expr Node
}

type ReturnStatement struct {
	Expr Node
}

func CreateExprStatement(expr Node) Node {
	return ExprStatement{expr}
}

func CreateReturnStatement(expr Node) Node {
	return ReturnStatement{expr}
}
