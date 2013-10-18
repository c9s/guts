
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
const T_EQUAL = 57354
const T_LT_EQUAL = 57355
const T_GT_EQUAL = 57356
const T_INDENT = 57357
const T_OUTDENT = 57358
const T_NEWLINE = 57359
const T_NEW = 57360
const T_CLONE = 57361
const T_IF = 57362
const T_IN = 57363
const T_ELSEIF = 57364
const T_ELSE = 57365
const T_FOR = 57366
const T_SAY = 57367
const T_SPACE = 57368
const T_ECHO = 57369
const T_FOREACH = 57370
const T_TRY = 57371
const T_CATCH = 57372
const T_CLASS = 57373
const T_IS = 57374
const T_DOES = 57375
const T_FUNCTION_PROTOTYPE = 57376
const T_FUNCTION_GLYPH = 57377
const T_RANGE_OPERATOR = 57378
const T_CONST = 57379
const T_RETURN = 57380
const T_BREAK = 57381
const T_CONTINUE = 57382
const T_THROW = 57383
const T_NS_SEPARATOR = 57384
const T_NAMESPACE = 57385
const T_CALLSTART = 57386
const T_CALLEND = 57387
const T_OBJECT_OPERATOR = 57388
const UMINUS = 57389
const T_BOOLEAN_OR = 57390
const T_BOOLEAN_AND = 57391

var GutsToknames = []string{
	"T_DOT",
	"T_IDENTIFIER",
	"T_FLOATING",
	"T_NUMBER",
	"T_STRING",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
	"T_EOF",
	"T_EQUAL",
	"T_LT_EQUAL",
	"T_GT_EQUAL",
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

//line src/gutscript/parser.y:378
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 56
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 241

var GutsAct = []int{

	7, 78, 69, 6, 87, 96, 81, 83, 79, 89,
	34, 35, 36, 38, 67, 41, 42, 34, 35, 36,
	86, 46, 80, 45, 75, 37, 49, 50, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 39, 23,
	22, 97, 72, 4, 65, 24, 70, 61, 62, 13,
	71, 77, 63, 24, 31, 3, 30, 25, 26, 27,
	28, 29, 73, 32, 33, 44, 74, 88, 48, 40,
	32, 33, 34, 35, 36, 13, 79, 84, 45, 85,
	14, 70, 47, 90, 93, 92, 15, 82, 94, 95,
	1, 34, 35, 36, 43, 91, 2, 11, 98, 34,
	35, 36, 10, 8, 16, 68, 76, 66, 9, 31,
	19, 30, 25, 26, 27, 28, 29, 34, 35, 36,
	18, 0, 34, 35, 36, 32, 33, 64, 31, 0,
	30, 25, 26, 27, 28, 29, 31, 0, 30, 25,
	26, 27, 28, 29, 32, 33, 34, 35, 36, 0,
	0, 0, 32, 33, 0, 0, 30, 25, 26, 27,
	28, 29, 25, 26, 27, 28, 29, 39, 23, 22,
	32, 33, 17, 23, 22, 32, 33, 0, 0, 0,
	0, 0, 13, 0, 5, 0, 0, 20, 27, 28,
	29, 21, 17, 23, 22, 0, 0, 0, 0, 32,
	33, 0, 13, 0, 0, 12, 0, 20, 0, 14,
	0, 21, 0, 0, 14, 15, 0, 0, 0, 0,
	15, 0, 0, 0, 0, 12, 0, 0, 0, 0,
	0, 0, 0, 0, 14, 0, 0, 0, 0, 0,
	15,
}
var GutsPact = []int{

	167, -1000, -1000, 28, -1000, -1000, -1000, 87, -1000, -1000,
	8, -1000, 162, 167, 162, 162, -1000, 31, -1000, -1000,
	162, 77, -1000, -1000, 187, 162, 162, 162, 162, 162,
	162, 162, 162, 162, 162, 162, 162, 25, 87, -24,
	36, 79, -2, 162, -33, 162, 60, 21, -1000, 134,
	134, -2, -2, -2, 110, 105, 87, 87, 87, 87,
	87, 162, 34, -1000, -1000, 87, -11, 3, -42, -1000,
	87, -1000, -54, 60, -1000, 34, -44, -1000, -1000, 4,
	162, -1000, -1000, 33, -1000, -1000, 71, -1000, 162, -1000,
	-1000, -1000, -57, 5, -1000, 87, -1000, 162, 87,
}
var GutsPgo = []int{

	0, 0, 120, 110, 43, 108, 107, 106, 1, 2,
	105, 104, 55, 103, 102, 97, 3, 96, 90, 87,
	85,
}
var GutsR1 = []int{

	0, 18, 17, 16, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 4, 19, 19, 15, 14, 14, 14,
	13, 8, 8, 8, 7, 7, 6, 6, 6, 5,
	20, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 9, 10, 10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 3, 4, 3, 5, 4,
	3, 3, 2, 1, 3, 1, 3, 2, 0, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 0, 4,
}
var GutsChk = []int{

	-1000, -18, -17, -12, -4, 17, -16, -1, -13, -5,
	-14, -15, 38, 15, 47, 53, -11, 5, -2, -3,
	20, 24, 7, 6, 17, 52, 53, 54, 55, 56,
	51, 49, 65, 66, 12, 13, 14, 17, -1, 5,
	-12, -1, -1, 63, 34, 47, -1, 5, -4, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, 22, 23, 16, 48, -1, -6, 47, -10, -9,
	-1, -16, 21, -1, -16, 35, -7, 48, -8, 5,
	64, 48, -19, 61, -16, -16, 64, 48, 63, 5,
	-9, 62, -20, -1, -8, -1, 62, 36, -1,
}
var GutsDef = []int{

	0, -2, 1, 2, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 50, 49, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 13, 46,
	0, 0, 44, 0, 28, 54, 0, 0, 4, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 0, 0, 3, 31, 20, 0, 0, 0, 53,
	51, 17, 0, 0, 19, 0, 0, 27, 25, 23,
	0, 55, 16, 0, 18, 29, 0, 26, 0, 22,
	52, 14, 0, 0, 24, 21, 15, 0, 30,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 57, 3, 3, 3, 56, 51, 3,
	47, 48, 54, 52, 64, 53, 3, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 63, 65, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 61, 3, 62, 50, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 49,
}
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 58, 59, 60,
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
		//line src/gutscript/parser.y:138
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:145
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:151
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:155
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:162
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 6:
		//line src/gutscript/parser.y:167
		{ }
	case 7:
		//line src/gutscript/parser.y:171
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:172
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:173
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:174
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:175
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:176
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:177
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 16:
		//line src/gutscript/parser.y:186
		{  }
	case 17:
		//line src/gutscript/parser.y:192
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementList))
	        }
	case 18:
		//line src/gutscript/parser.y:197
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 19:
		//line src/gutscript/parser.y:203
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 20:
		//line src/gutscript/parser.y:211
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 21:
		//line src/gutscript/parser.y:219
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 22:
		//line src/gutscript/parser.y:223
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 23:
		//line src/gutscript/parser.y:227
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 24:
		//line src/gutscript/parser.y:233
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 25:
		//line src/gutscript/parser.y:239
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 26:
		//line src/gutscript/parser.y:245
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 27:
		//line src/gutscript/parser.y:248
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 28:
		//line src/gutscript/parser.y:251
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 29:
		//line src/gutscript/parser.y:257
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 31:
		//line src/gutscript/parser.y:266
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
		//line src/gutscript/parser.y:276
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 33:
		//line src/gutscript/parser.y:280
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 34:
		//line src/gutscript/parser.y:284
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 35:
		//line src/gutscript/parser.y:288
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 36:
		//line src/gutscript/parser.y:292
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 37:
		//line src/gutscript/parser.y:296
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 38:
		//line src/gutscript/parser.y:300
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 39:
		//line src/gutscript/parser.y:304
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 40:
		//line src/gutscript/parser.y:308
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 41:
		//line src/gutscript/parser.y:312
		{
	            GutsVAL.val = ast.CreateExpr(T_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 42:
		//line src/gutscript/parser.y:316
		{
	            GutsVAL.val = ast.CreateExpr(T_LT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 43:
		//line src/gutscript/parser.y:320
		{
	            GutsVAL.val = ast.CreateExpr(T_GT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 44:
		//line src/gutscript/parser.y:324
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 45:
		//line src/gutscript/parser.y:327
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 46:
		//line src/gutscript/parser.y:330
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 47:
		//line src/gutscript/parser.y:334
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 48:
		//line src/gutscript/parser.y:337
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 49:
		//line src/gutscript/parser.y:342
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 50:
		//line src/gutscript/parser.y:346
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 51:
		//line src/gutscript/parser.y:350
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 52:
		//line src/gutscript/parser.y:355
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 53:
		//line src/gutscript/parser.y:362
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 54:
		//line src/gutscript/parser.y:367
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 55:
		//line src/gutscript/parser.y:373
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
