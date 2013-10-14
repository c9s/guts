%{
package gutscript

// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "gutscript/ast"

var regs = make([]int, 26)
var base int

const DEBUG = false

func debug(msg string, vals ...interface{}) {
    if DEBUG {
        fmt.Print(msg)
        for _, val := range vals {
            fmt.Printf(" %#v",val)
        }
        fmt.Println("\n")
    }
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
%type <val> function_parameter_list
%type <val> function_parameters
%type <val> function_parameter
%type <val> function_call_parameter
%type <val> function_call_parameter_list
%type <val> function_call
%type <val> statement 
%type <val> statement_list
%type <val> assignment_statement 
%type <val> if_statement
%type <val> block 
%type <val> top_statement 
%type <val> top_statement_list 
%type <val> start

// same for terminals
%token <val> T_DOT T_IDENTIFIER T_FLOATING T_NUMBER T_STRING

%token T_ONELINE_COMMENT T_COMMENT
%token T_EOF

%token T_INDENT T_OUTDENT

%token T_NEWLINE

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
%token T_FUNCTION_GLYPH
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


// block is contains one or more statement
block: T_INDENT statement_list T_OUTDENT { $$ = $2 }

statement_list:
      statement_list statement
      {
            if stmts, ok := $1.(*ast.StatementNodeList) ; ok {
                stmts.Append($2)
                $$ = $1
            }
      }
    | statement
      { 
            stmts := ast.CreateStatementList()
            stmts.Append($1)
            $$ = stmts
      }
;

statement: 
          block { $$ = $1 }
        | expr { $$ = ast.CreateExprStatement($1) } 
        | assignment_statement { $$ = $1 }
        | function_decl_statement { $$ = $1 }
        | if_statement { $$ = $1 }
        | T_RETURN expr { $$ = ast.CreateReturnStatement($2) }
    ;

if_statement:
        T_IF expr block
        {
            $$ = ast.CreateIfStatement($2.(ast.Expr), $3.(*ast.StatementNodeList))
        }
    |
        if_statement T_ELSEIF expr block 
        {
            $1.(*ast.IfStatement).AddElseIf($3.(ast.Expr),$4.(*ast.StatementNodeList))
            $$ = $1
        }
    | 
        if_statement T_ELSE block
        {
            $1.(*ast.IfStatement).SetElse($3.(*ast.StatementNodeList))
            $$ = $1
        }
;

assignment_statement:
    variable '=' expr
        {
            debug("assignment_statement", $1 , "=" , $3)
            $$ = ast.CreateAssignStatement($1, $3)
        }
    | variable '=' expr T_NEWLINE
        {
            debug("assignment_statement", $1 , "=" , $3)
            $$ = ast.CreateAssignStatement($1, $3)
        }
;


function_parameter: 
    /*
    T_IDENTIFIER '=' expr {
        $$ = ast.FunctionParam{$1.(string), "", $3}
    }
    |
    */
    T_IDENTIFIER T_IDENTIFIER {
        $$ = ast.FunctionParam{$2.(string), $1.(string), nil}
    }
    | 
    T_IDENTIFIER { 
        $$ = ast.FunctionParam{$1.(string), "", nil}
    }
    ;

function_parameters: 
      function_parameters ',' function_parameter {
        if params, ok := $1.([]ast.FunctionParam) ; ok {
            params = append(params, $3.(ast.FunctionParam))
            $$ = params
        }
      }
    | function_parameter {
            $$ = []ast.FunctionParam{$1.(ast.FunctionParam)}
        }
    ;

function_parameter_list: 
    '(' function_parameters ')' {
        $$ = $2
    }
    | '(' ')' {
        $$ = []ast.FunctionParam{}
    }
    | /* empty */ { 
        $$ = []ast.FunctionParam{}
    }
    ;

function_decl_statement:
    T_IDENTIFIER T_FUNCTION_PROTOTYPE function_parameter_list T_FUNCTION_GLYPH block { 
        $$ = ast.CreateFunction($1.(string), $3.([]ast.FunctionParam), $5.(*ast.StatementNodeList))
    }
;


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
    | function_call { 
            $$ = ast.UnaryExpr{0, $1}
        }
    | variable { 
            $$ = ast.UnaryExpr{0, $1}
        }
    | number {
            $$ = ast.UnaryExpr{0, $1}
        }
    | floating_number {
            $$ = ast.UnaryExpr{0, $1}
        }
    ;

floating_number: T_FLOATING {
        $$ = ast.CreateFloatingNumber($1.(string))
    }

number: T_NUMBER {
        $$ = ast.CreateNumber($1.(string))
    }

function_call_parameter: expr {
        $$ = ast.UnaryExpr{0, $1}
    };

function_call_parameter_list:
    function_call_parameter_list ',' function_call_parameter {
        if params, ok := $1.([]ast.Node) ; ok {
            params = append(params, $3.(ast.Node))
            $$ = params
        }
    }
    | 
    function_call_parameter {
        // create the expr list
        $$ = []ast.Node{$1}
    }
;

function_call:
    T_IDENTIFIER '(' function_call_parameter_list ')' {
        $$ = ast.CreateFunctionCall($1.(string), $3.([]ast.Node))
    }
;

%%      /*  start  of  programs  */
