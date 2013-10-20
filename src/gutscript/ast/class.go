package ast

type Class struct {
	Name string
	// Properties []Node
	// Methods    []Node
	Super      *string
	Interfaces []string
	Body       *StatementList
}

func CreateClass(className string) Node {
	return Class{className, nil, []string{}, nil}
}

func (self Class) SetSuper(className string) {
	self.Super = &className
}

func (self Class) AddInterface(inf string) {
	self.Interfaces = append(self.Interfaces, inf)
}

func (self Class) SetInterfaces(infs []string) {
	self.Interfaces = infs
}
