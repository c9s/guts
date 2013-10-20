package ast

type Function struct {
	Name   string
	Params []FunctionParam
	Body   *StatementList
}

type FunctionParam struct {
	Name    string
	Type    string
	Default Node
}

func CreateFunction(name string, params []FunctionParam, stmts *StatementList) Function {
	return Function{name, params, stmts}
}
