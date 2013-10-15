
//line src/gutscript/parser.y:2
package gutscript
import __yyfmt__ "fmt"
//line src/gutscript/parser.y:2
		
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


//line src/gutscript/parser.y:26
type GutsSymType struct{
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
const T_INDENT = 57354
const T_OUTDENT = 57355
const T_NEWLINE = 57356
const T_NEW = 57357
const T_CLONE = 57358
const T_IF = 57359
const T_ELSEIF = 57360
const T_ELSE = 57361
const T_FOR = 57362
const T_SAY = 57363
const T_SPACE = 57364
const T_ECHO = 57365
const T_FOREACH = 57366
const T_TRY = 57367
const T_CATCH = 57368
const T_CLASS = 57369
const T_IS = 57370
const T_DOES = 57371
const T_FUNCTION_PROTOTYPE = 57372
const T_FUNCTION_GLYPH = 57373
const T_RANGE_OPERATOR = 57374
const T_CONST = 57375
const T_RETURN = 57376
const T_BREAK = 57377
const T_CONTINUE = 57378
const T_THROW = 57379
const T_NS_SEPARATOR = 57380
const T_NAMESPACE = 57381
const T_CALLSTART = 57382
const T_CALLEND = 57383
const T_OBJECT_OPERATOR = 57384
const UMINUS = 57385
const T_BOOLEAN_OR = 57386
const T_BOOLEAN_AND = 57387

var GutsToknames = []string{
	"T_DOT",
	"T_IDENTIFIER",
	"T_FLOATING",
	"T_NUMBER",
	"T_STRING",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
	"T_EOF",
	"T_INDENT",
	"T_OUTDENT",
	"T_NEWLINE",
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
	"T_FUNCTION_GLYPH",
	"T_RANGE_OPERATOR",
	"T_CONST",
	"T_RETURN",
	"T_BREAK",
	"T_CONTINUE",
	"T_THROW",
	"T_NS_SEPARATOR",
	"T_NAMESPACE",
	"T_CALLSTART",
	"T_CALLEND",
	"T_OBJECT_OPERATOR",
	" (",
	" )",
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
var GutsStatenames = []string{}

const GutsEofCode = 1
const GutsErrCode = 2
const GutsMaxDepth = 200

//line src/gutscript/parser.y:353
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 50
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 173

var GutsAct = []int{

	7, 68, 60, 30, 31, 75, 6, 12, 71, 58,
	39, 77, 33, 69, 36, 37, 40, 65, 32, 74,
	41, 22, 70, 40, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 4, 12, 38, 3, 56,
	29, 61, 28, 23, 24, 25, 26, 27, 62, 54,
	22, 35, 67, 63, 69, 1, 30, 31, 42, 2,
	64, 55, 29, 76, 28, 23, 24, 25, 26, 27,
	72, 61, 73, 78, 10, 8, 79, 80, 30, 31,
	29, 15, 28, 23, 24, 25, 26, 27, 28, 23,
	24, 25, 26, 27, 59, 66, 30, 31, 25, 26,
	27, 57, 30, 31, 23, 24, 25, 26, 27, 30,
	31, 9, 16, 21, 20, 18, 17, 30, 31, 12,
	0, 5, 0, 0, 19, 16, 21, 20, 34, 21,
	20, 0, 12, 0, 0, 0, 0, 19, 0, 0,
	0, 11, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 11, 0, 14, 0, 0, 0,
	0, 0, 0, 13, 0, 0, 13, 0, 0, 14,
	0, 0, 14,
}
var GutsPact = []int{

	107, -1000, -1000, 7, -1000, -1000, -1000, 35, -1000, -1000,
	4, 123, 107, 123, 123, -1000, -20, -1000, -1000, 123,
	-1000, -1000, 120, 123, 123, 123, 123, 123, 123, 123,
	123, 123, 15, 35, -27, 36, 17, -58, 123, -34,
	123, -5, -1000, 48, 48, -58, -58, -58, 56, 41,
	35, 35, 123, 24, -1000, -1000, 35, -14, 8, -36,
	-1000, 35, -1000, -5, -1000, 24, -39, -1000, -1000, 6,
	123, -1000, -1000, -1000, 49, -1000, 123, -1000, -1000, -1000,
	35,
}
var GutsPgo = []int{

	0, 0, 116, 115, 35, 111, 101, 95, 1, 2,
	94, 81, 38, 75, 74, 6, 59, 55, 55, 55,
}
var GutsR1 = []int{

	0, 17, 16, 15, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 14, 14, 14, 13, 8, 8, 8,
	7, 7, 6, 6, 6, 5, 18, 19, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 2, 9, 10, 10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 2, 3, 5, 4, 3, 3, 2, 1,
	3, 1, 3, 2, 0, 5, 1, 5, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 0, 4,
}
var GutsChk = []int{

	-1000, -17, -16, -12, -4, 14, -15, -1, -13, -5,
	-14, 34, 12, 43, 49, -11, 5, -2, -3, 17,
	7, 6, 14, 48, 49, 50, 51, 52, 47, 45,
	61, 62, 14, -1, 5, -12, -1, -1, 57, 30,
	43, -1, -4, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, 18, 19, 13, 44, -1, -6, 43, -10,
	-9, -1, -15, -1, -15, 31, -7, 44, -8, 5,
	58, 44, -15, -15, 58, 44, 57, 5, -9, -8,
	-1,
}
var GutsDef = []int{

	0, -2, 1, 2, 5, 6, 7, 8, 9, 10,
	11, 0, 0, 0, 0, 39, 40, 41, 42, 0,
	44, 43, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 40, 0, 0, 38, 0, 24,
	48, 0, 4, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 0, 0, 3, 28, 16, 0, 0, 0,
	47, 45, 13, 0, 15, 0, 0, 23, 21, 19,
	0, 49, 14, 25, 0, 22, 0, 18, 46, 20,
	17,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 3, 3, 3, 52, 47, 3,
	43, 44, 50, 48, 58, 49, 3, 51, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	62, 57, 61, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 59, 3, 60, 46, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 45,
}
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 54, 55, 56,
}
var GutsTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var GutsDebug = 100

type GutsLexer interface {
	Lex(lval *GutsSymType) int
	Error(s string)
}

const GutsFlag = -1000

func GutsTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(GutsToknames) {
		if GutsToknames[c-4] != "" {
			return GutsToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func GutsStatname(s int) string {
	if s >= 0 && s < len(GutsStatenames) {
		if GutsStatenames[s] != "" {
			return GutsStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Gutslex1(lex GutsLexer, lval *GutsSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = GutsTok1[0]
		goto out
	}
	if char < len(GutsTok1) {
		c = GutsTok1[char]
		goto out
	}
	if char >= GutsPrivate {
		if char < GutsPrivate+len(GutsTok2) {
			c = GutsTok2[char-GutsPrivate]
			goto out
		}
	}
	for i := 0; i < len(GutsTok3); i += 2 {
		c = GutsTok3[i+0]
		if c == char {
			c = GutsTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = GutsTok2[1] /* unknown char */
	}
	if GutsDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", GutsTokname(c), uint(char))
	}
	return c
}

type GutsParser struct {
	Val GutsSymType
}

func (self *GutsParser) Parse(Gutslex GutsLexer) int {
	var Gutsn int
	var Gutslval GutsSymType
	var GutsVAL GutsSymType
	GutsS := make([]GutsSymType, GutsMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Gutsstate := 0
	Gutschar := -1
	Gutsp := -1
	goto Gutsstack

ret0:
	return 0

ret1:
	return 1

Gutsstack:
	/* put a state and value onto the stack */
	if GutsDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", GutsTokname(Gutschar), GutsStatname(Gutsstate))
	}

	Gutsp++
	if Gutsp >= len(GutsS) {
		nyys := make([]GutsSymType, len(GutsS)*2)
		copy(nyys, GutsS)
		GutsS = nyys
	}
	GutsS[Gutsp] = GutsVAL
	GutsS[Gutsp].yys = Gutsstate

Gutsnewstate:
	Gutsn = GutsPact[Gutsstate]
	if Gutsn <= GutsFlag {
		goto Gutsdefault /* simple state */
	}
	if Gutschar < 0 {
		Gutschar = Gutslex1(Gutslex, &Gutslval)
	}
	Gutsn += Gutschar
	if Gutsn < 0 || Gutsn >= GutsLast {
		goto Gutsdefault
	}
	Gutsn = GutsAct[Gutsn]
	if GutsChk[Gutsn] == Gutschar { /* valid shift */
		Gutschar = -1
		GutsVAL = Gutslval
		Gutsstate = Gutsn
		if Errflag > 0 {
			Errflag--
		}
		goto Gutsstack
	}

Gutsdefault:
	/* default state action */
	Gutsn = GutsDef[Gutsstate]
	if Gutsn == -2 {
		if Gutschar < 0 {
			Gutschar = Gutslex1(Gutslex, &Gutslval)
		}

		/* look through exception table */
		xi := 0
		for {
			if GutsExca[xi+0] == -1 && GutsExca[xi+1] == Gutsstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Gutsn = GutsExca[xi+0]
			if Gutsn < 0 || Gutsn == Gutschar {
				break
			}
		}
		Gutsn = GutsExca[xi+1]
		if Gutsn < 0 {
			goto ret0
		}
	}
	if Gutsn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Gutslex.Error("syntax error")
			Nerrs++
			if GutsDebug >= 1 {
				__yyfmt__.Printf("%s", GutsStatname(Gutsstate))
				__yyfmt__.Printf(" saw %s\n", GutsTokname(Gutschar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Gutsp >= 0 {
				Gutsn = GutsPact[GutsS[Gutsp].yys] + GutsErrCode
				if Gutsn >= 0 && Gutsn < GutsLast {
					Gutsstate = GutsAct[Gutsn] /* simulate a shift of "error" */
					if GutsChk[Gutsstate] == GutsErrCode {
						goto Gutsstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if GutsDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", GutsS[Gutsp].yys)
				}
				Gutsp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if GutsDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", GutsTokname(Gutschar))
			}
			if Gutschar == GutsEofCode {
				goto ret1
			}
			Gutschar = -1
			goto Gutsnewstate /* try again in the same state */
		}
	}

	/* reduction by production Gutsn */
	if GutsDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Gutsn, GutsStatname(Gutsstate))
	}

	Gutsnt := Gutsn
	Gutspt := Gutsp
	_ = Gutspt // guard against "declared and not used"

	Gutsp -= GutsR2[Gutsn]
	GutsVAL = GutsS[Gutsp+1]

	self.Val = GutsVAL

	/* consult goto table to find next state */
	Gutsn = GutsR1[Gutsn]
	Gutsg := GutsPgo[Gutsn]
	Gutsj := Gutsg + GutsS[Gutsp].yys + 1

	if Gutsj >= GutsLast {
		Gutsstate = GutsAct[Gutsg]
	} else {
		Gutsstate = GutsAct[Gutsj]
		if GutsChk[Gutsstate] != -Gutsn {
			Gutsstate = GutsAct[Gutsg]
		}
	}
	// dummy call; replaced with literal code
	switch Gutsnt {

	case 1:
		//line src/gutscript/parser.y:133
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:140
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:146
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:150
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:157
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 6:
		//line src/gutscript/parser.y:162
		{ }
	case 7:
		//line src/gutscript/parser.y:166
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:167
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:168
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:169
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:170
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:171
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 13:
		//line src/gutscript/parser.y:176
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementList))
	        }
	case 14:
		//line src/gutscript/parser.y:181
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 15:
		//line src/gutscript/parser.y:187
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 16:
		//line src/gutscript/parser.y:195
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 17:
		//line src/gutscript/parser.y:203
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 18:
		//line src/gutscript/parser.y:207
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 19:
		//line src/gutscript/parser.y:211
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 20:
		//line src/gutscript/parser.y:217
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 21:
		//line src/gutscript/parser.y:223
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 22:
		//line src/gutscript/parser.y:229
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 23:
		//line src/gutscript/parser.y:232
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 24:
		//line src/gutscript/parser.y:235
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 25:
		//line src/gutscript/parser.y:241
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 28:
		//line src/gutscript/parser.y:253
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 29:
		//line src/gutscript/parser.y:263
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 30:
		//line src/gutscript/parser.y:267
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 31:
		//line src/gutscript/parser.y:271
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 32:
		//line src/gutscript/parser.y:275
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 33:
		//line src/gutscript/parser.y:279
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 34:
		//line src/gutscript/parser.y:283
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 35:
		//line src/gutscript/parser.y:287
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 36:
		//line src/gutscript/parser.y:291
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 37:
		//line src/gutscript/parser.y:295
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 38:
		//line src/gutscript/parser.y:299
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 39:
		//line src/gutscript/parser.y:302
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 40:
		//line src/gutscript/parser.y:305
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 41:
		//line src/gutscript/parser.y:309
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 42:
		//line src/gutscript/parser.y:312
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 43:
		//line src/gutscript/parser.y:317
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 44:
		//line src/gutscript/parser.y:321
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 45:
		//line src/gutscript/parser.y:325
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 46:
		//line src/gutscript/parser.y:330
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 47:
		//line src/gutscript/parser.y:337
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 48:
		//line src/gutscript/parser.y:342
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 49:
		//line src/gutscript/parser.y:348
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
