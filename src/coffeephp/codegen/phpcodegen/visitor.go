package phpcodegen

import "coffeephp/ast"
import "strconv"
import "fmt"

type Visitor struct{}

func Visit(n ast.Node) string {
	// fmt.Printf("visit %#v\n", n)
	if stmts, ok := n.(*ast.StatementNodeList); ok {
		var output string
		for _, stmt := range *stmts {
			output += Visit(stmt)
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
		return fmt.Sprintf("%c%s", expr.Op, Visit(expr.Val))
	}
	if expr, ok := n.(ast.Expr); ok {
		if expr.Parenthesis {
			return fmt.Sprintf("(%s %c %s)", Visit(expr.Left), expr.Op, Visit(expr.Right))
		}
		return fmt.Sprintf("%s %c %s", Visit(expr.Left), expr.Op, Visit(expr.Right))
	}
	if stmt, ok := n.(ast.AssignStatement); ok {
		return Visit(stmt.Variable) + " = " + Visit(stmt.Expr) + ";\n"
	}
	return ""
}
