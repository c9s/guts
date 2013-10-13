package phpcodegen

import "coffeephp/ast"
import "strconv"
import "strings"
import "fmt"

type Visitor struct {
	indent int
}

func (self *Visitor) IndentSpace() string {
	if self.indent == 0 {
		return ""
	}
	return strings.Repeat("    ", self.indent)
}

func (self *Visitor) Visit(n ast.Node) string {
	// fmt.Printf("visit %#v\n", n)
	if stmts, ok := n.(*ast.StatementNodeList); ok {
		var output string
		for _, stmt := range *stmts {
			output += self.IndentSpace() + self.Visit(stmt)
		}
		return output
	}
	if variable, ok := n.(ast.VariableNode); ok {
		return "$" + variable.Identifier
	}
	if number, ok := n.(ast.NumberNode); ok {
		return strconv.FormatInt(number.Val, 10)
	}
	if floating, ok := n.(ast.FloatingNumberNode); ok {
		return strconv.FormatFloat(floating.Val, 'e', -1, 64)
	}
	if expr, ok := n.(ast.UnaryExpr); ok {
		return fmt.Sprintf("%c%s", expr.Op, self.Visit(expr.Val))
	}
	if expr, ok := n.(ast.Expr); ok {
		if expr.Parenthesis {
			return fmt.Sprintf("(%s %c %s)", self.Visit(expr.Left), expr.Op, self.Visit(expr.Right))
		}
		return fmt.Sprintf("%s %c %s", self.Visit(expr.Left), expr.Op, self.Visit(expr.Right))
	}
	if stmt, ok := n.(ast.AssignStatement); ok {
		return self.Visit(stmt.Variable) + " = " + self.Visit(stmt.Expr) + ";\n"
	}
	if stmt, ok := n.(ast.IfStatement); ok {
		var out string = ""
		out += self.IndentSpace() + "if ( " + self.Visit(stmt.Expr) + " ) {\n"
		self.indent++
		out += self.Visit(stmt.Body)
		self.indent--
		out += self.IndentSpace() + "\n}"
		return out
	}
	return ""
}
