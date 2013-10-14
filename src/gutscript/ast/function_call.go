package ast

type FunctionCall struct {
	Name   string
	Params []Expr
}

func CreateFunctionCall(name string, params []Expr) Node {
	return FunctionCall{name, params}
}
