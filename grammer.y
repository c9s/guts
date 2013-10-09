%{

package main

import _ "fmt"

var regs = make([]int, 26)
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number

// same for terminals
%token <val> T_DIGIT T_LETTER T_DOT T_IDENTIFIER EOF

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
%token T_ECHO
%token T_FOREACH
%token T_TRY
%token T_CATCH
%token T_CLASS "class (T_CLASS)"
%token T_IS   "is (T_IS)"
%token T_DOES "does (T_DOES)"
%token T_FUNCTION_PROTOTYPE ":: (T_FUNCTION_PROTOTYPE)"

%token T_CONST      "const (T_CONST)"
%token T_RETURN     "return (T_RETURN)"

%token T_BREAK      "break (T_BREAK)"
%token T_CONTINUE   "continue (T_CONTINUE)"

%token T_THROW      "throw (T_THROW)"
%token T_NS_SEPARATOR    "\\ (T_NS_SEPARATOR)"
%token T_NAMESPACE       "namespace (T_NAMESPACE)"


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

start : top_statement_list	{ }

top_statement_list:
		top_statement_list  { } top_statement { }
	|	/* empty */
;

top_statement:
    statement	{ }
;

statement:
     unticked_statement { }
   | assign_statement { }
;

unticked_statement:
	expr ';'				{ }
;

assign_statement:
      identity '=' expr ';' {  }
    | identity '=' function_call ';' {  }
;

function_parameter_list: '(' ')' ;

function:
      identity T_FUNCTION_PROTOTYPE function_parameter_list '->' function_body
;

function_body: top_statement_list;

expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$  =  $1 + $3 }
	|    expr '-' expr
		{ $$  =  $1 - $3 }
	|    expr '*' expr
		{ $$  =  $1 * $3 }
	|    expr '/' expr
		{ $$  =  $1 / $3 }
	|    expr '%' expr
		{ $$  =  $1 % $3 }
	|    expr '&' expr
		{ $$  =  $1 & $3 }
	|    expr '|' expr
		{ $$  =  $1 | $3 }
	|    '-'  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    T_LETTER
		{ $$  = regs[$1] }
	|    number
	;


identity: T_LETTER
        | T_LETTER T_DIGIT
        ;

number	:    T_DIGIT
		{
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}
	|    number T_DIGIT
		{ $$ = base * $1 + $2 }
	;

function_call_parameter_list:
    '(' ')'	{ }
;

function_call:
    identity function_call_parameter_list { }
;

%%      /*  start  of  programs  */
