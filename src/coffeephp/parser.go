
//line src/coffeephp/parser.y:2
package coffeephp
import __yyfmt__ "fmt"
//line src/coffeephp/parser.y:2
		
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


//line src/coffeephp/parser.y:26
type CoffeeSymType struct{
	yys int
    typ TokenType
    val ast.Node
    line int
    pos  int
}

const T_DOT = 57346
const T_IDENTIFIER = 57347
const T_FLOATING = 57348
const T_NUMBER = 57349
const T_STRING = 57350
const T_ONELINE_COMMENT = 57351
const T_COMMENT = 57352
const T_EOF = 57353
const T_INDENT_ENTER = 57354
const T_INDENT_EXIT = 57355
const T_NEWLINE = 57356
const T_ASSIGN = 57357
const T_NEW = 57358
const T_CLONE = 57359
const T_IF = 57360
const T_ELSEIF = 57361
const T_ELSE = 57362
const T_FOR = 57363
const T_SAY = 57364
const T_SPACE = 57365
const T_ECHO = 57366
const T_FOREACH = 57367
const T_TRY = 57368
const T_CATCH = 57369
const T_CLASS = 57370
const T_IS = 57371
const T_DOES = 57372
const T_FUNCTION_PROTOTYPE = 57373
const T_FUNCTION_ARROW = 57374
const T_RANGE_OPERATOR = 57375
const T_CONST = 57376
const T_RETURN = 57377
const T_BREAK = 57378
const T_CONTINUE = 57379
const T_THROW = 57380
const T_NS_SEPARATOR = 57381
const T_NAMESPACE = 57382
const T_OBJECT_OPERATOR = 57383
const UMINUS = 57384
const T_BOOLEAN_OR = 57385
const T_BOOLEAN_AND = 57386

var CoffeeToknames = []string{
	"T_DOT",
	"T_IDENTIFIER",
	"T_FLOATING",
	"T_NUMBER",
	"T_STRING",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
	"T_EOF",
	"T_INDENT_ENTER",
	"T_INDENT_EXIT",
	"T_NEWLINE",
	"T_ASSIGN",
	"T_NEW",
	"T_CLONE",
	"T_IF",
	"T_ELSEIF",
	"T_ELSE",
	"T_FOR",
	"T_SAY",
	"T_SPACE",
	"T_ECHO",
	"T_FOREACH",
	"T_TRY",
	"T_CATCH",
	"T_CLASS",
	"T_IS",
	"T_DOES",
	"T_FUNCTION_PROTOTYPE",
	"T_FUNCTION_ARROW",
	"T_RANGE_OPERATOR",
	"T_CONST",
	"T_RETURN",
	"T_BREAK",
	"T_CONTINUE",
	"T_THROW",
	"T_NS_SEPARATOR",
	"T_NAMESPACE",
	"T_OBJECT_OPERATOR",
	" |",
	" ^",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	" !",
	"UMINUS",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
}
var CoffeeStatenames = []string{}

const CoffeeEofCode = 1
const CoffeeErrCode = 2
const CoffeeMaxDepth = 200

//line src/coffeephp/parser.y:268
      /*  start  of  programs  */

//line yacctab:1
var CoffeeExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const CoffeeNprod = 35
const CoffeePrivate = 57344

var CoffeeTokenNames []string
var CoffeeStates []string

const CoffeeLast = 140

var CoffeeAct = []int{

	27, 49, 26, 21, 22, 23, 24, 25, 21, 22,
	23, 24, 25, 50, 28, 29, 23, 24, 25, 28,
	29, 33, 10, 52, 10, 28, 29, 28, 29, 27,
	2, 26, 21, 22, 23, 24, 25, 3, 48, 51,
	19, 31, 30, 28, 29, 27, 20, 26, 21, 22,
	23, 24, 25, 26, 21, 22, 23, 24, 25, 28,
	29, 11, 18, 17, 53, 28, 29, 54, 47, 57,
	11, 18, 17, 10, 12, 10, 1, 10, 10, 34,
	18, 17, 55, 12, 56, 8, 6, 5, 7, 4,
	16, 15, 0, 19, 19, 0, 0, 0, 0, 0,
	0, 0, 14, 0, 0, 0, 0, 0, 9, 0,
	13, 14, 0, 0, 0, 0, 0, 0, 0, 13,
	14, 32, 35, 36, 0, 0, 0, 0, 13, 0,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
}
var CoffeePact = []int{

	65, -1000, 65, -1000, 32, -1000, -1000, -1000, -1000, 3,
	27, 10, 74, 74, 74, -1000, -1000, -1000, -1000, -1000,
	-1000, 74, 74, 74, 74, 74, 74, 74, 74, 74,
	74, -16, -13, -1000, -1000, -42, -29, -31, -31, -29,
	-29, -29, -37, 9, 3, 3, 3, 7, -32, 52,
	-1000, 65, -1000, 65, -1000, 65, 56, -1000,
}
var CoffeePgo = []int{

	0, 108, 21, 91, 90, 89, 88, 87, 86, 85,
	37, 30, 76, 68, 67, 67, 67,
}
var CoffeeR1 = []int{

	0, 12, 11, 11, 10, 10, 5, 5, 5, 5,
	7, 8, 9, 13, 6, 14, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 3, 15, 16,
}
var CoffeeR2 = []int{

	0, 1, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 3, 6, 2, 5, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 1, 1,
	1, 1, 1, 2, 2,
}
var CoffeeChk = []int{

	-1000, -12, -11, -10, -5, -7, -8, -6, -9, -1,
	-2, 5, 18, 54, 46, -3, -4, 7, 6, -10,
	14, 45, 46, 47, 48, 49, 44, 42, 56, 57,
	15, 31, -1, -2, 5, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -13, 54, 14,
	55, 32, 55, 12, -14, -11, -11, 13,
}
var CoffeeDef = []int{

	0, -2, 1, 3, 4, 6, 7, 8, 9, 10,
	28, 16, 0, 0, 0, 29, 30, 32, 31, 2,
	5, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 28, 16, 0, 27, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 11, 0, 0, 0,
	17, 0, 13, 0, 14, 15, 0, 12,
}
var CoffeeTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 3, 3, 3, 49, 44, 3,
	54, 55, 47, 45, 3, 46, 3, 48, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	57, 3, 56, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 43, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 42,
}
var CoffeeTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	51, 52, 53,
}
var CoffeeTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var CoffeeDebug = 10

type CoffeeLexer interface {
	Lex(lval *CoffeeSymType) int
	Error(s string)
}

const CoffeeFlag = -1000

func CoffeeTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 57346 && c-57346 < len(CoffeeToknames) {
		if CoffeeToknames[c-57346] != "" {
			return CoffeeToknames[c-57346]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func CoffeeStatname(s int) string {
	if s >= 0 && s < len(CoffeeStatenames) {
		if CoffeeStatenames[s] != "" {
			return CoffeeStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Coffeelex1(lex CoffeeLexer, lval *CoffeeSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = CoffeeTok1[0]
		goto out
	}
	if char < len(CoffeeTok1) {
		c = CoffeeTok1[char]
		goto out
	}
	if char >= CoffeePrivate {
		if char < CoffeePrivate+len(CoffeeTok2) {
			c = CoffeeTok2[char-CoffeePrivate]
			goto out
		}
	}
	for i := 0; i < len(CoffeeTok3); i += 2 {
		c = CoffeeTok3[i+0]
		if c == char {
			c = CoffeeTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = CoffeeTok2[1] /* unknown char */
	}
	if CoffeeDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", CoffeeTokname(c), uint(char))
	}
	return c
}

type CoffeeParser struct {
	Val CoffeeSymType
}

func (self *CoffeeParser) Parse(Coffeelex CoffeeLexer) int {
	var Coffeen int
	var Coffeelval CoffeeSymType
	var CoffeeVAL CoffeeSymType
	CoffeeS := make([]CoffeeSymType, CoffeeMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Coffeestate := 0
	Coffeechar := -1
	Coffeep := -1
	goto Coffeestack

ret0:
	return 0

ret1:
	return 1

Coffeestack:
	/* put a state and value onto the stack */
	if CoffeeDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", CoffeeTokname(Coffeechar), CoffeeStatname(Coffeestate))
	}

	Coffeep++
	if Coffeep >= len(CoffeeS) {
		nyys := make([]CoffeeSymType, len(CoffeeS)*2)
		copy(nyys, CoffeeS)
		CoffeeS = nyys
	}
	CoffeeS[Coffeep] = CoffeeVAL
	CoffeeS[Coffeep].yys = Coffeestate

Coffeenewstate:
	Coffeen = CoffeePact[Coffeestate]
	if Coffeen <= CoffeeFlag {
		goto Coffeedefault /* simple state */
	}
	if Coffeechar < 0 {
		Coffeechar = Coffeelex1(Coffeelex, &Coffeelval)
	}
	Coffeen += Coffeechar
	if Coffeen < 0 || Coffeen >= CoffeeLast {
		goto Coffeedefault
	}
	Coffeen = CoffeeAct[Coffeen]
	if CoffeeChk[Coffeen] == Coffeechar { /* valid shift */
		Coffeechar = -1
		CoffeeVAL = Coffeelval
		Coffeestate = Coffeen
		if Errflag > 0 {
			Errflag--
		}
		goto Coffeestack
	}

Coffeedefault:
	/* default state action */
	Coffeen = CoffeeDef[Coffeestate]
	if Coffeen == -2 {
		if Coffeechar < 0 {
			Coffeechar = Coffeelex1(Coffeelex, &Coffeelval)
		}

		/* look through exception table */
		xi := 0
		for {
			if CoffeeExca[xi+0] == -1 && CoffeeExca[xi+1] == Coffeestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Coffeen = CoffeeExca[xi+0]
			if Coffeen < 0 || Coffeen == Coffeechar {
				break
			}
		}
		Coffeen = CoffeeExca[xi+1]
		if Coffeen < 0 {
			goto ret0
		}
	}
	if Coffeen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Coffeelex.Error("syntax error")
			Nerrs++
			if CoffeeDebug >= 1 {
				__yyfmt__.Printf("%s", CoffeeStatname(Coffeestate))
				__yyfmt__.Printf(" saw %s\n", CoffeeTokname(Coffeechar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Coffeep >= 0 {
				Coffeen = CoffeePact[CoffeeS[Coffeep].yys] + CoffeeErrCode
				if Coffeen >= 0 && Coffeen < CoffeeLast {
					Coffeestate = CoffeeAct[Coffeen] /* simulate a shift of "error" */
					if CoffeeChk[Coffeestate] == CoffeeErrCode {
						goto Coffeestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if CoffeeDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", CoffeeS[Coffeep].yys)
				}
				Coffeep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if CoffeeDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", CoffeeTokname(Coffeechar))
			}
			if Coffeechar == CoffeeEofCode {
				goto ret1
			}
			Coffeechar = -1
			goto Coffeenewstate /* try again in the same state */
		}
	}

	/* reduction by production Coffeen */
	if CoffeeDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Coffeen, CoffeeStatname(Coffeestate))
	}

	Coffeent := Coffeen
	Coffeept := Coffeep
	_ = Coffeept // guard against "declared and not used"

	Coffeep -= CoffeeR2[Coffeen]
	CoffeeVAL = CoffeeS[Coffeep+1]

	self.Val = CoffeeVAL

	/* consult goto table to find next state */
	Coffeen = CoffeeR1[Coffeen]
	Coffeeg := CoffeePgo[Coffeen]
	Coffeej := Coffeeg + CoffeeS[Coffeep].yys + 1

	if Coffeej >= CoffeeLast {
		Coffeestate = CoffeeAct[Coffeeg]
	} else {
		Coffeestate = CoffeeAct[Coffeej]
		if CoffeeChk[Coffeestate] != -Coffeen {
			Coffeestate = CoffeeAct[Coffeeg]
		}
	}
	// dummy call; replaced with literal code
	switch Coffeent {

	case 1:
		//line src/coffeephp/parser.y:122
		{
	        debug("end compilation", CoffeeS[Coffeept-0].val)
	        CoffeeVAL.val = CoffeeS[Coffeept-0].val
	    }
	case 2:
		//line src/coffeephp/parser.y:129
		{ 
	        if stmts, ok := CoffeeS[Coffeept-1].val.(*ast.StatementNodeList) ; ok {
	            stmts.Append(CoffeeS[Coffeept-0].val)
	            CoffeeVAL.val = CoffeeS[Coffeept-1].val
	        }
	        debug("top_statement_list top_statement", CoffeeS[Coffeept-1].val, CoffeeS[Coffeept-0].val)
	    }
	case 3:
		//line src/coffeephp/parser.y:136
		{
	        debug("top_statement one", CoffeeS[Coffeept-0].val)
	        stmts := ast.CreateStatementNodeList()
	        stmts.Append(CoffeeS[Coffeept-0].val)
	        CoffeeVAL.val = stmts
	    }
	case 4:
		//line src/coffeephp/parser.y:145
		{
	        debug("statement", CoffeeS[Coffeept-0].val)
	        CoffeeVAL.val = CoffeeS[Coffeept-0].val
	    }
	case 5:
		//line src/coffeephp/parser.y:149
		{
	        debug("statement.newline", CoffeeS[Coffeept-1].val)
	        CoffeeVAL.val = CoffeeS[Coffeept-1].val
	    }
	case 6:
		//line src/coffeephp/parser.y:156
		{ CoffeeVAL.val = CoffeeS[Coffeept-0].val }
	case 7:
		//line src/coffeephp/parser.y:157
		{ CoffeeVAL.val = CoffeeS[Coffeept-0].val }
	case 8:
		//line src/coffeephp/parser.y:158
		{ CoffeeVAL.val = CoffeeS[Coffeept-0].val }
	case 9:
		//line src/coffeephp/parser.y:159
		{ CoffeeVAL.val = CoffeeS[Coffeept-0].val }
	case 10:
		//line src/coffeephp/parser.y:162
		{ CoffeeVAL.val = ast.CreateExprStatement(CoffeeS[Coffeept-0].val) }
	case 11:
		//line src/coffeephp/parser.y:166
		{
	        debug("assignment_statement", CoffeeS[Coffeept-2].val , "=" , CoffeeS[Coffeept-0].val)
	        CoffeeVAL.val = ast.CreateAssignStatement(CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	    }
	case 12:
		//line src/coffeephp/parser.y:173
		{
	        CoffeeVAL.val = ast.CreateIfStatement(CoffeeS[Coffeept-4].val.(ast.Expr), CoffeeS[Coffeept-1].val.(*ast.StatementNodeList))
	    }
	case 14:
		//line src/coffeephp/parser.y:180
		{ 
	        // $1 ($3) $4
    }
	case 16:
		//line src/coffeephp/parser.y:193
		{ 
	        CoffeeVAL.val = ast.CreateVariableNode(CoffeeS[Coffeept-0].val.(string))
	    }
	case 17:
		//line src/coffeephp/parser.y:198
		{
	            if node, ok := CoffeeS[Coffeept-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                CoffeeVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 18:
		//line src/coffeephp/parser.y:208
		{ 
	            CoffeeVAL.val = ast.CreateExpr('+', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 19:
		//line src/coffeephp/parser.y:212
		{ 
	            CoffeeVAL.val = ast.CreateExpr('-', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 20:
		//line src/coffeephp/parser.y:216
		{ 
	            CoffeeVAL.val = ast.CreateExpr('*', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 21:
		//line src/coffeephp/parser.y:220
		{ 
	            CoffeeVAL.val = ast.CreateExpr('/', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 22:
		//line src/coffeephp/parser.y:224
		{ 
	            CoffeeVAL.val = ast.CreateExpr('%', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 23:
		//line src/coffeephp/parser.y:228
		{
	            CoffeeVAL.val = ast.CreateExpr('&', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 24:
		//line src/coffeephp/parser.y:232
		{
	            CoffeeVAL.val = ast.CreateExpr('|', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 25:
		//line src/coffeephp/parser.y:236
		{
	            CoffeeVAL.val = ast.CreateExpr('>', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 26:
		//line src/coffeephp/parser.y:240
		{
	            CoffeeVAL.val = ast.CreateExpr('<', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 27:
		//line src/coffeephp/parser.y:244
		{ 
	            CoffeeVAL.val = ast.UnaryExpr{'-', CoffeeS[Coffeept-0].val}
	        }
	case 28:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 29:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 30:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 31:
		//line src/coffeephp/parser.y:252
		{
	        CoffeeVAL.val = ast.CreateFloatingNumberNode(CoffeeS[Coffeept-0].val.(string))
	    }
	case 32:
		//line src/coffeephp/parser.y:256
		{
	        CoffeeVAL.val = ast.CreateNumberNode(CoffeeS[Coffeept-0].val.(string))
	    }
	case 33:
		//line src/coffeephp/parser.y:261
		{ }
	case 34:
		//line src/coffeephp/parser.y:265
		{ }
	}
	goto Coffeestack /* stack new state and value */
}
