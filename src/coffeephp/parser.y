%{
package coffeephp

// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "coffeephp/ast"

var regs = make([]int, 26)
var base int

const DEBUG = true

func debug(msg string, vals ...interface{}) {
    fmt.Print(msg)
    for _, val := range vals {
        fmt.Printf(" %#v",val)
    }
    fmt.Println("\n")
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
    typ TokenType
    val ast.Node
    line int
    pos  int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr variable number floating_number
%type <val> statement
%type <val> function_decl_statement
%type <val> unticked_statement 
%type <val> assignment_statement 
%type <val> inner_statement 
%type <val> inner_statement_list
%type <val> top_statement 
%type <val> top_statement_list 
%type <val> start

// same for terminals
%token <val> T_DOT T_IDENTIFIER T_FLOATING T_NUMBER T_STRING

%token T_ONELINE_COMMENT T_COMMENT
%token T_EOF

%token T_INDENT_ENTER T_OUTDENT

%token T_NEWLINE

%token T_ASSIGN

%token T_NEW
%token T_CLONE

%token T_IF

%left T_ELSEIF
%token T_ELSEIF

%left T_ELSE 
%token T_ELSE

%token T_FOR

// T_SAY is basically an alias of T_ECHO
%token T_SAY
%token T_SPACE
%token T_ECHO
%token T_FOREACH
%token T_TRY
%token T_CATCH
%token T_CLASS
%token T_IS
%token T_DOES

// ::
%token T_FUNCTION_PROTOTYPE 

// ->
%token T_FUNCTION_ARROW
%token T_RANGE_OPERATOR

%token T_STRING

%token T_CONST
%token T_RETURN

%token T_BREAK
%token T_CONTINUE

%token T_THROW
%token T_NS_SEPARATOR
%token T_NAMESPACE

// obj.method
%token T_OBJECT_OPERATOR

%left '|'
%left '^'
%left '&'

%left '+' '-'
%left '*'  '/'  '%'
%right '!'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%left T_BOOLEAN_OR
%token T_BOOLEAN_OR
%left T_BOOLEAN_AND 
%token T_BOOLEAN_AND


%start start

%%

start : top_statement_list {
        debug("end compilation", $1)
        $$ = $1
    }
;

top_statement_list:
    top_statement_list top_statement { 
        if stmts, ok := $1.(*ast.StatementNodeList) ; ok {
            stmts.Append($2)
            $$ = $1
        }
        debug("top_statement_list top_statement", $1, $2)
    }
    | top_statement {
        debug("top_statement", $1)
        stmts := ast.CreateStatementList()
        stmts.Append($1)
        $$ = stmts
    }
;

top_statement:
    statement {
        debug("statement", $1)
        $$ = $1
    }
;

inner_statement: statement;

inner_statement_list:
      inner_statement_list inner_statement 
      {
            if stmts, ok := $1.(*ast.StatementNodeList) ; ok {
                stmts.Append($2)
                $$ = $1
            }
      }
    | inner_statement 
      { 
            stmts := ast.CreateStatementList()
            stmts.Append($1)
            $$ = stmts
      }
;

statement:
      unticked_statement { $$ = $1 }
;

unticked_statement: 
          T_NEWLINE T_INDENT_ENTER inner_statement_list T_NEWLINE T_OUTDENT { $$ = $3 }
        | expr { $$ = ast.CreateExprStatement($1) } 
        | assignment_statement { $$ = $1 }
        | function_decl_statement { $$ = $1 }
        | T_IF expr statement T_ELSE statement
          {
                debug("if-else-statement")
                $$ = ast.CreateIfElseStatement($2.(ast.Expr), $3.(*ast.StatementNodeList), $5.(*ast.StatementNodeList))
          }
        | T_IF expr statement 
          {
                debug("if-statement")
                $$ = ast.CreateIfStatement($2.(ast.Expr), $3.(*ast.StatementNodeList))
          }
    ;


assignment_statement:
        variable T_ASSIGN expr 
        {
            debug("assignment_statement", $1 , "=" , $3)
            $$ = ast.CreateAssignStatement($1, $3)
        }
      | variable T_ASSIGN expr T_NEWLINE 
        {
            debug("assignment_statement", $1 , "=" , $3)
            $$ = ast.CreateAssignStatement($1, $3)
        }
;

function_parameter_list: '(' ')' ;

function_decl_statement:
    T_IDENTIFIER T_FUNCTION_PROTOTYPE function_parameter_list T_FUNCTION_ARROW function_body { 
        // $1 ($3) $4
    }
;

function_body: top_statement_list;







variable: T_IDENTIFIER { 
        $$ = ast.CreateVariable($1.(string))
    }

expr:
      '(' expr ')' {
            if node, ok := $2.(ast.Expr) ; ok {
                node.Parenthesis = true
                $$ = node
            } else {
                panic(" type casting to ast.Expr failed.")
            }
            // $$ = $2
        }
    | expr '+' expr
        { 
            $$ = ast.CreateExpr('+', $1, $3)
        }
    | expr '-' expr
        { 
            $$ = ast.CreateExpr('-', $1, $3)
        }
    | expr '*' expr
        { 
            $$ = ast.CreateExpr('*', $1, $3)
        }
    | expr '/' expr
        { 
            $$ = ast.CreateExpr('/', $1, $3)
        }
    | expr '%' expr 
        { 
            $$ = ast.CreateExpr('%', $1, $3)
        }
    | expr '&' expr 
        {
            $$ = ast.CreateExpr('&', $1, $3)
        }
    | expr '|' expr 
        {
            $$ = ast.CreateExpr('|', $1, $3)
        }
    | expr '>' expr 
        {
            $$ = ast.CreateExpr('>', $1, $3)
        }
    | expr '<' expr 
        {
            $$ = ast.CreateExpr('<', $1, $3)
        }
    | '-' expr  %prec UMINUS
        { 
            $$ = ast.UnaryExpr{'-', $2}
        }
    | variable
    | number
    | floating_number
    ;

floating_number: T_FLOATING {
        $$ = ast.CreateFloatingNumber($1.(string))
    }

number: T_NUMBER {
        $$ = ast.CreateNumber($1.(string))
    }

function_call_parameter_list:
    '(' ')' { }
;

function_call:
    T_IDENTIFIER function_call_parameter_list { }
;

%%      /*  start  of  programs  */
