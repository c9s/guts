%{
package gutscript

// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "gutscript/ast"

var regs = make([]int, 26)
var base int

func debug(msg string, vals ...interface{}) {
    if GutsDebug > 0 {
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
%type <val> expr number floating_number
%type <val> statement
%type <val> function_decl_statement
%type <val> function_parameter_list
%type <val> function_parameters
%type <val> function_parameter
%type <val> function_call_parameter
%type <val> function_call_parameter_list
%type <val> function_call
%type <val> statement_list
%type <val> assignment
%type <val> if_statement
%type <val> for_statement
%type <val> block 
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
%token T_IN

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

%token T_CALLSTART
%token T_CALLEND

// obj.method
%token T_OBJECT_OPERATOR

// for function call
%left '(' ')'
%left T_CALLSTART T_CALLEND

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

top_statement_list: statement_list 
    { 
        $$ = $1 
    }
    ;

// block is contains one or more statement
block: T_INDENT statement_list T_OUTDENT { $$ = $2 }

statement_list:
      statement_list T_NEWLINE statement
      {
            if stmts, ok := $1.(*ast.StatementList) ; ok {
                stmts.Append($3)
                $$ = $1
            }
      }
    | statement
      { 
            stmts := ast.CreateStatementList()
            stmts.Append($1)
            $$ = stmts
      }
    | T_NEWLINE { }
;

statement: 
          block { $$ = $1 }
        | expr { $$ = ast.CreateExprStatement($1) } 
        | assignment { $$ = $1 }
        | function_decl_statement { $$ = $1 }
        | if_statement { $$ = $1 }
        | for_statement { $$ = $1 }
        | T_RETURN expr { $$ = ast.CreateReturnStatement($2) }
    ;

listop: 
      '['  ']'
      | '[' range ']'
      ;

for_statement:
    T_FOR T_IDENTIFIER T_IN listop {  }
    ;


if_statement:
        T_IF expr block
        {
            $$ = ast.CreateIfStatement($2.(ast.Expr), $3.(*ast.StatementList))
        }
    |
        if_statement T_NEWLINE T_ELSEIF expr block 
        {
            $1.(*ast.IfStatement).AddElseIf($4.(ast.Expr),$5.(*ast.StatementList))
            $$ = $1
        }
    | 
        if_statement T_NEWLINE T_ELSE block
        {
            $1.(*ast.IfStatement).SetElse($4.(*ast.StatementList))
            $$ = $1
        }
;

assignment:
     T_IDENTIFIER '=' expr
        {
            debug("assignment", $1 , "=" , $3)
            $$ = ast.CreateAssignStatement(ast.CreateVariable($1.(string)), $3)
        }
;


function_parameter: 
    T_IDENTIFIER '=' expr {
        $$ = ast.FunctionParam{$1.(string), "", $3}
    }
    |
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
        debug("function declaration", $1, $3, $5)
        $$ = ast.CreateFunction($1.(string), $3.([]ast.FunctionParam), $5.(*ast.StatementList))
    }
;

range: expr T_RANGE_OPERATOR expr ;

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
    | T_IDENTIFIER { 
            // $$ = ast.UnaryExpr{0, $1}
            $$ = ast.CreateVariable($1.(string)) 
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
    |
    /* */ {
        $$ = []ast.Node{}
    }
;

function_call:
    T_IDENTIFIER '(' function_call_parameter_list ')' {
        $$ = ast.CreateFunctionCall($1.(string), $3.([]ast.Node))
    }
;

%%      /*  start  of  programs  */
