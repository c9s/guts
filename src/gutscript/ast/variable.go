package ast

type Variable struct {
	Identifier string
}

func CreateVariable(identifier string) Node {
	return Variable{Identifier: identifier}
}
