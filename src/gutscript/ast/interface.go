package ast

type Interface struct {
	Name string
}

func CreateInterface(name string) Node {
	return Interface{name}
}
