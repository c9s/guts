package ast

type Function struct {
	Name   string
	Params []FunctionParam
	Body   *StatementNodeList
}

type FunctionParam struct {
	Name string
	Type string
}

func CreateFunction(name string, params []FunctionParam, stmts *StatementNodeList) Node {
	return Function{name, params, stmts}
}
