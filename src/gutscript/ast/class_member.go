package ast

type ClassMember struct {
	Name  string
	Value Node
}

func CreateClassMember(name string) ClassMember {
	return ClassMember{name, nil}
}

func (self *ClassMember) SetValue(val Node) {
	self.Value = &val
}
