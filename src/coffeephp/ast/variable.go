package ast

type VariableNode struct {
	Identifier string
}

func CreateVariableNode(identifier string) Node {
	return VariableNode{Identifier: identifier}
}
