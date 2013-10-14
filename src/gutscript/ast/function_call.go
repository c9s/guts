package ast

type FunctionCall struct {
	Name   string
	Params []Node
}

func CreateFunctionCall(name string, params []Node) Node {
	return FunctionCall{name, params}
}
