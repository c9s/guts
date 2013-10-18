package gutscript

import "fmt"

func (self *GutsSymType) GetLine() int {
	return self.line
}

func (self *GutsSymType) GetPos() int {
	return self.pos
}

func (self *GutsSymType) String() string {
	switch self.typ {
	case T_EOF, eof:
		return "EOF"
	}
	return fmt.Sprintf("Token: %s %q", GetTokenName(int(self.typ)), self.val)
}
