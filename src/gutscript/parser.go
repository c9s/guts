
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
const T_IN = 57360
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
const T_FUNCTION_GLYPH = 57374
const T_RANGE_OPERATOR = 57375
const T_CONST = 57376
const T_RETURN = 57377
const T_BREAK = 57378
const T_CONTINUE = 57379
const T_THROW = 57380
const T_NS_SEPARATOR = 57381
const T_NAMESPACE = 57382
const T_CALLSTART = 57383
const T_CALLEND = 57384
const T_OBJECT_OPERATOR = 57385
const UMINUS = 57386
const T_BOOLEAN_OR = 57387
const T_BOOLEAN_AND = 57388

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
	"T_IN",
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

//line src/gutscript/parser.y:362
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 53
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 196

var GutsAct = []int{

	7, 72, 63, 6, 30, 25, 26, 27, 28, 29,
	32, 33, 90, 35, 61, 38, 39, 83, 32, 33,
	81, 43, 75, 77, 42, 69, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 91, 80, 4, 74, 55,
	56, 59, 66, 64, 73, 41, 34, 65, 31, 24,
	30, 25, 26, 27, 28, 29, 67, 13, 42, 13,
	68, 73, 45, 44, 32, 33, 25, 26, 27, 28,
	29, 78, 82, 79, 40, 64, 86, 84, 87, 32,
	33, 76, 88, 89, 71, 1, 57, 24, 27, 28,
	29, 31, 92, 30, 25, 26, 27, 28, 29, 32,
	33, 36, 23, 22, 2, 3, 11, 32, 33, 58,
	31, 10, 30, 25, 26, 27, 28, 29, 31, 37,
	30, 25, 26, 27, 28, 29, 32, 33, 8, 17,
	23, 22, 16, 62, 32, 33, 13, 70, 5, 60,
	14, 20, 17, 23, 22, 21, 15, 9, 19, 13,
	36, 23, 22, 18, 20, 85, 0, 0, 21, 12,
	0, 0, 0, 0, 0, 0, 0, 0, 14, 0,
	0, 0, 12, 0, 15, 0, 0, 0, 0, 0,
	0, 14, 0, 0, 0, 0, 0, 15, 0, 14,
	0, 0, 0, 0, 0, 15,
}
var GutsPact = []int{

	124, -1000, -1000, 35, -1000, -1000, -1000, 72, -1000, -1000,
	32, -1000, 145, 124, 145, 145, -1000, 14, -1000, -1000,
	145, 58, -1000, -1000, 137, 145, 145, 145, 145, 145,
	145, 145, 145, 145, 20, 72, -20, 73, 64, -52,
	145, -30, 145, 45, 24, -1000, 37, 37, -52, -52,
	-52, 17, -44, 72, 72, 145, 47, -1000, -1000, 72,
	-7, 39, -23, -1000, 72, -1000, -35, 45, -1000, 47,
	-25, -1000, -1000, 12, 145, -1000, -1000, 96, -1000, -1000,
	56, -1000, 145, -1000, -1000, -1000, -47, 2, -1000, 72,
	-1000, 145, 72,
}
var GutsPgo = []int{

	0, 0, 153, 148, 37, 147, 139, 137, 1, 2,
	133, 132, 105, 128, 111, 106, 3, 104, 85, 81,
	76,
}
var GutsR1 = []int{

	0, 18, 17, 16, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 4, 19, 19, 15, 14, 14, 14,
	13, 8, 8, 8, 7, 7, 6, 6, 6, 5,
	20, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 9, 10,
	10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 3, 4, 3, 5, 4,
	3, 3, 2, 1, 3, 1, 3, 2, 0, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 0, 4,
}
var GutsChk = []int{

	-1000, -18, -17, -12, -4, 14, -16, -1, -13, -5,
	-14, -15, 35, 12, 44, 50, -11, 5, -2, -3,
	17, 21, 7, 6, 14, 49, 50, 51, 52, 53,
	48, 46, 62, 63, 14, -1, 5, -12, -1, -1,
	60, 31, 44, -1, 5, -4, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, 19, 20, 13, 45, -1,
	-6, 44, -10, -9, -1, -16, 18, -1, -16, 32,
	-7, 45, -8, 5, 61, 45, -19, 58, -16, -16,
	61, 45, 60, 5, -9, 59, -20, -1, -8, -1,
	59, 33, -1,
}
var GutsDef = []int{

	0, -2, 1, 2, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 47, 46, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 13, 43, 0, 0, 41,
	0, 28, 51, 0, 0, 4, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 0, 0, 3, 31, 20,
	0, 0, 0, 50, 48, 17, 0, 0, 19, 0,
	0, 27, 25, 23, 0, 52, 16, 0, 18, 29,
	0, 26, 0, 22, 49, 14, 0, 0, 24, 21,
	15, 0, 30,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 54, 3, 3, 3, 53, 48, 3,
	44, 45, 51, 49, 61, 50, 3, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	63, 60, 62, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 58, 3, 59, 47, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 46,
}
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 55, 56, 57,
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
		//line src/gutscript/parser.y:134
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:141
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:147
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:151
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:158
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 6:
		//line src/gutscript/parser.y:163
		{ }
	case 7:
		//line src/gutscript/parser.y:167
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:168
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:169
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:170
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:171
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:172
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:173
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 16:
		//line src/gutscript/parser.y:182
		{  }
	case 17:
		//line src/gutscript/parser.y:188
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementList))
	        }
	case 18:
		//line src/gutscript/parser.y:193
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 19:
		//line src/gutscript/parser.y:199
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 20:
		//line src/gutscript/parser.y:207
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 21:
		//line src/gutscript/parser.y:215
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 22:
		//line src/gutscript/parser.y:219
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 23:
		//line src/gutscript/parser.y:223
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 24:
		//line src/gutscript/parser.y:229
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 25:
		//line src/gutscript/parser.y:235
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 26:
		//line src/gutscript/parser.y:241
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 27:
		//line src/gutscript/parser.y:244
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 28:
		//line src/gutscript/parser.y:247
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 29:
		//line src/gutscript/parser.y:253
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 31:
		//line src/gutscript/parser.y:262
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 32:
		//line src/gutscript/parser.y:272
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 33:
		//line src/gutscript/parser.y:276
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 34:
		//line src/gutscript/parser.y:280
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 35:
		//line src/gutscript/parser.y:284
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 36:
		//line src/gutscript/parser.y:288
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 37:
		//line src/gutscript/parser.y:292
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 38:
		//line src/gutscript/parser.y:296
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 39:
		//line src/gutscript/parser.y:300
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 40:
		//line src/gutscript/parser.y:304
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 41:
		//line src/gutscript/parser.y:308
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 42:
		//line src/gutscript/parser.y:311
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 43:
		//line src/gutscript/parser.y:314
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 44:
		//line src/gutscript/parser.y:318
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 45:
		//line src/gutscript/parser.y:321
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 46:
		//line src/gutscript/parser.y:326
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 47:
		//line src/gutscript/parser.y:330
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 48:
		//line src/gutscript/parser.y:334
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 49:
		//line src/gutscript/parser.y:339
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 50:
		//line src/gutscript/parser.y:346
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 51:
		//line src/gutscript/parser.y:351
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 52:
		//line src/gutscript/parser.y:357
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
