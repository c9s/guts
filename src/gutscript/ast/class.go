package ast

type ClassStatement struct {
	Name       string
	Properties []Node
	Methods    []Node
	Super      *string
	Interfaces []string
}

func CreateClassStatement(className string) Node {
	return ClassStatement{className, nil, nil, nil, []string{}}
}

func (self ClassStatement) SetSuperClass(className string) {
	self.Super = &className
}

func (self ClassStatement) AddInterface(intf string) {
	self.Interfaces = append(self.Interfaces, intf)
}
