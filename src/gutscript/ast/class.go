package ast

type ClassStatement struct {
	Name       string
	Properties []Node
	Methods    []Node
}

func CreateClassStatement(className string) Node {
	return ClassStatement{className, nil, nil}
}
