%{
package coffeephp

// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "coffeephp/ast"

var regs = make([]int, 26)
var base int

var Stmts []ast.Node
var Namespaces []ast.Node

const DEBUG = true

func debug(msg string, vals ...interface{}) {
    fmt.Print(msg, ": ")
    fmt.Println(vals...)
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
%type <val> expr statement variable
%type <val> unticked_statement assignment_statement top_statement

// same for terminals
%token <val> T_DIGIT T_LETTER T_DOT T_IDENTIFIER T_FLOATING T_NUMBER T_STRING

%token T_ONELINE_COMMENT T_COMMENT
%token T_EOF

%token T_INDENT_ENTER T_INDENT_EXIT

%token T_PLUS

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
%token T_FUNCTION_PROTOTYPE
%token T_RANGE_OPERATOR
%token T_BRACKET_OPEN T_BRACKET_CLOSE

%token T_STRING

%token T_CONST
%token T_RETURN

%token T_BREAK
%token T_CONTINUE

%token T_THROW
%token T_NS_SEPARATOR
%token T_NAMESPACE


// obj.method
%token T_OBJECT_OPERATOR ". (T_OBJECT_OPERATOR)"

%left 'and'
%left 'or'

%left '|'
%left '^'
%left '&'

%left '+'  '-'
%left '*'  '/'  '%'
%right '!'

%left T_BOOLEAN_OR
%token T_BOOLEAN_OR
%left T_BOOLEAN_AND 
%token T_BOOLEAN_AND

%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%start start

%%

start : top_statement_list {
        debug("end compilation")
    }
;

top_statement_list:
        top_statement_list {
            debug("top_statement_list")
        }
        top_statement {
            debug("top_statement_list.top_statement")
        }
    |   /* empty */
;

top_statement:
    statement {
        debug("top_statement done", $1)
        Stmts = append(Stmts, $1)
    }
    | statement T_NEWLINE {
        debug("top_statement.NEWLINE done", $1)
        Stmts = append(Stmts, $1)
    }
;

statement:
      unticked_statement { $$ = $1 }
    | assignment_statement { $$ = $1 }
;

unticked_statement: expr { $$ = ast.CreateExprStatementNode($1) } ;

assignment_statement:
    variable T_ASSIGN expr {
        debug("assignment_statement", $1 , "=" , $3)
        $$ = ast.CreateAssignStatementNode($1, $3)
    }
;

variable: T_IDENTIFIER

function_parameter_list: '(' ')' ;

function:
      T_IDENTIFIER T_FUNCTION_PROTOTYPE function_parameter_list '->' function_body
;

function_body: top_statement_list;

expr    :
      '(' expr ')' {
            fmt.Println("wrap expr")
        }
    | expr '+' expr
        { 
            $$ = ast.CreateExprNode('+', $1, $3)
        }
    | expr '-' expr
        { 
            $$ = ast.CreateExprNode('-', $1, $3)
        }
    | expr '*' expr
        { 
            $$ = ast.CreateExprNode('*', $1, $3)
        }
    | expr '/' expr
        { 
            $$ = ast.CreateExprNode('/', $1, $3)
        }
    | expr '%' expr
        { 
            $$ = ast.CreateExprNode('%', $1, $3)
        }
    | expr '&' expr
        { 
            $$ = ast.CreateExprNode('&', $1, $3)
        }
    | expr '|' expr
        { 
            $$ = ast.CreateExprNode('|', $1, $3)
        }
    | '-' expr  %prec UMINUS
        { 
            // $$  = -$2  
        }
    | T_IDENTIFIER
        { 
            // $$  = regs[$1] 
        }
    | T_NUMBER {
        $$ = ast.CreateNumberNode($1.(string))
    }
    | T_FLOATING {
        $$ = ast.CreateFloatingNumberNode($1.(string))
    }
    ;

function_call_parameter_list:
    '(' ')' { }
;

function_call:
    T_IDENTIFIER function_call_parameter_list { }
;

%%      /*  start  of  programs  */
