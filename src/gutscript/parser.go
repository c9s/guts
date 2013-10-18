
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
const T_EXTENDS = 57376
const T_FUNCTION_PROTOTYPE = 57377
const T_FUNCTION_GLYPH = 57378
const T_RANGE_OPERATOR = 57379
const T_CONST = 57380
const T_RETURN = 57381
const T_BREAK = 57382
const T_CONTINUE = 57383
const T_THROW = 57384
const T_NS_SEPARATOR = 57385
const T_NAMESPACE = 57386
const T_CALLSTART = 57387
const T_CALLEND = 57388
const T_OBJECT_OPERATOR = 57389
const UMINUS = 57390
const T_BOOLEAN_OR = 57391
const T_BOOLEAN_AND = 57392

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
	"T_EXTENDS",
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

//line src/gutscript/parser.y:403
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 65
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 248

var GutsAct = []int{

	7, 83, 72, 6, 106, 95, 41, 25, 24, 97,
	107, 86, 91, 84, 40, 70, 43, 44, 36, 37,
	38, 94, 48, 47, 36, 37, 38, 85, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	88, 80, 77, 108, 76, 109, 68, 39, 73, 15,
	64, 65, 74, 4, 46, 16, 33, 82, 32, 27,
	28, 29, 30, 31, 101, 78, 26, 47, 96, 79,
	66, 26, 34, 35, 36, 37, 38, 14, 34, 35,
	51, 14, 92, 45, 93, 3, 73, 84, 98, 100,
	89, 50, 103, 36, 37, 38, 104, 105, 49, 99,
	42, 36, 37, 38, 87, 75, 102, 90, 1, 110,
	2, 12, 33, 11, 32, 27, 28, 29, 30, 31,
	36, 37, 38, 10, 8, 36, 37, 38, 34, 35,
	67, 33, 17, 32, 27, 28, 29, 30, 31, 33,
	71, 32, 27, 28, 29, 30, 31, 34, 35, 36,
	37, 38, 41, 25, 24, 34, 35, 81, 69, 9,
	32, 27, 28, 29, 30, 31, 27, 28, 29, 30,
	31, 18, 25, 24, 34, 35, 20, 19, 0, 34,
	35, 14, 0, 5, 0, 0, 21, 0, 0, 0,
	23, 0, 29, 30, 31, 15, 0, 22, 18, 25,
	24, 16, 0, 34, 35, 13, 0, 0, 14, 0,
	0, 0, 0, 21, 15, 0, 0, 23, 0, 0,
	16, 0, 0, 0, 22, 0, 0, 0, 0, 0,
	0, 0, 13, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 0, 0, 0, 0, 16,
}
var GutsPact = []int{

	166, -1000, -1000, 49, -1000, -1000, -1000, 89, -1000, -1000,
	30, -1000, -1000, 147, 166, 147, 147, -1000, 19, -1000,
	-1000, 147, 93, 86, -1000, -1000, 193, 147, 147, 147,
	147, 147, 147, 147, 147, 147, 147, 147, 147, 28,
	89, -25, 54, 81, 12, 147, -33, 147, 62, 10,
	21, -1000, 137, 137, 12, 12, 12, 113, 108, 89,
	89, 89, 89, 89, 147, 66, -1000, -1000, 89, 5,
	8, -38, -1000, 89, -1000, 7, 85, -50, 62, -1000,
	66, -44, -1000, -1000, 4, 147, -1000, -1000, 84, -1000,
	-1000, 1, -1000, -1000, 82, -1000, 147, -1000, -1000, -61,
	-1000, -1000, -53, 6, -1000, 89, 40, -1000, 147, -1000,
	89,
}
var GutsPgo = []int{

	0, 0, 177, 176, 53, 159, 158, 157, 1, 2,
	140, 132, 85, 124, 123, 113, 111, 3, 110, 108,
	107, 106, 105, 104, 99,
}
var GutsR1 = []int{

	0, 19, 18, 17, 12, 12, 12, 4, 4, 4,
	4, 4, 4, 4, 4, 20, 20, 16, 14, 14,
	14, 13, 8, 8, 8, 7, 7, 6, 6, 6,
	5, 15, 22, 22, 23, 23, 24, 24, 24, 21,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 2,
	9, 10, 10, 10, 11,
}
var GutsR2 = []int{

	0, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 3, 4, 3, 5,
	4, 3, 3, 2, 1, 3, 1, 3, 2, 0,
	5, 4, 2, 0, 2, 0, 0, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 0, 4,
}
var GutsChk = []int{

	-1000, -19, -18, -12, -4, 17, -17, -1, -13, -5,
	-14, -15, -16, 39, 15, 48, 54, -11, 5, -2,
	-3, 20, 31, 24, 7, 6, 17, 53, 54, 55,
	56, 57, 52, 50, 66, 67, 12, 13, 14, 17,
	-1, 5, -12, -1, -1, 64, 35, 48, -1, 5,
	5, -4, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, 22, 23, 16, 49, -1, -6,
	48, -10, -9, -1, -17, -22, 34, 21, -1, -17,
	36, -7, 49, -8, 5, 65, 49, -23, 33, 5,
	-20, 62, -17, -17, 65, 49, 64, 5, -9, -24,
	5, 63, -21, -1, -8, -1, 65, 63, 37, 5,
	-1,
}
var GutsDef = []int{

	0, -2, 1, 2, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 0, 0, 0, 0, 54, 55, 56,
	57, 0, 0, 0, 59, 58, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 55, 0, 0, 53, 0, 29, 63, 0, 33,
	0, 4, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 0, 0, 3, 40, 21, 0,
	0, 0, 62, 60, 18, 35, 0, 0, 0, 20,
	0, 0, 28, 26, 24, 0, 64, 31, 36, 32,
	17, 0, 19, 30, 0, 27, 0, 23, 61, 34,
	38, 15, 0, 0, 25, 22, 0, 16, 0, 37,
	39,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 3, 3, 3, 57, 52, 3,
	48, 49, 55, 53, 65, 54, 3, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	67, 64, 66, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 63, 51, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 50,
}
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 59, 60, 61,
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
		//line src/gutscript/parser.y:140
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:147
		{ 
	        GutsVAL.val = GutsS[Gutspt-0].val 
	    }
	case 3:
		//line src/gutscript/parser.y:153
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 4:
		//line src/gutscript/parser.y:157
		{
	            if stmts, ok := GutsS[Gutspt-2].val.(*ast.StatementList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-2].val
	            }
	      }
	case 5:
		//line src/gutscript/parser.y:164
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 6:
		//line src/gutscript/parser.y:169
		{ }
	case 7:
		//line src/gutscript/parser.y:173
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 8:
		//line src/gutscript/parser.y:174
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 9:
		//line src/gutscript/parser.y:175
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 10:
		//line src/gutscript/parser.y:176
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:177
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:178
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:179
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 14:
		//line src/gutscript/parser.y:180
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 17:
		//line src/gutscript/parser.y:189
		{  }
	case 18:
		//line src/gutscript/parser.y:195
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementList))
	        }
	case 19:
		//line src/gutscript/parser.y:200
		{
	            GutsS[Gutspt-4].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-4].val
	        }
	case 20:
		//line src/gutscript/parser.y:206
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 21:
		//line src/gutscript/parser.y:214
		{
	            debug("assignment", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(ast.CreateVariable(GutsS[Gutspt-2].val.(string)), GutsS[Gutspt-0].val)
	        }
	case 22:
		//line src/gutscript/parser.y:222
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-2].val.(string), "", GutsS[Gutspt-0].val}
	    }
	case 23:
		//line src/gutscript/parser.y:226
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 24:
		//line src/gutscript/parser.y:230
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 25:
		//line src/gutscript/parser.y:236
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 26:
		//line src/gutscript/parser.y:242
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 27:
		//line src/gutscript/parser.y:248
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 28:
		//line src/gutscript/parser.y:251
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 29:
		//line src/gutscript/parser.y:254
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 30:
		//line src/gutscript/parser.y:260
		{ 
	        debug("function declaration", GutsS[Gutspt-4].val, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementList))
	    }
	case 31:
		//line src/gutscript/parser.y:267
		{  }
	case 32:
		//line src/gutscript/parser.y:271
		{  }
	case 34:
		//line src/gutscript/parser.y:276
		{  }
	case 37:
		//line src/gutscript/parser.y:281
		{ }
	case 38:
		//line src/gutscript/parser.y:282
		{  }
	case 40:
		//line src/gutscript/parser.y:291
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 41:
		//line src/gutscript/parser.y:301
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 42:
		//line src/gutscript/parser.y:305
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 43:
		//line src/gutscript/parser.y:309
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 44:
		//line src/gutscript/parser.y:313
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 45:
		//line src/gutscript/parser.y:317
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 46:
		//line src/gutscript/parser.y:321
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 47:
		//line src/gutscript/parser.y:325
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 48:
		//line src/gutscript/parser.y:329
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 49:
		//line src/gutscript/parser.y:333
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 50:
		//line src/gutscript/parser.y:337
		{
	            GutsVAL.val = ast.CreateExpr(T_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 51:
		//line src/gutscript/parser.y:341
		{
	            GutsVAL.val = ast.CreateExpr(T_LT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 52:
		//line src/gutscript/parser.y:345
		{
	            GutsVAL.val = ast.CreateExpr(T_GT_EQUAL, GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 53:
		//line src/gutscript/parser.y:349
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 54:
		//line src/gutscript/parser.y:352
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 55:
		//line src/gutscript/parser.y:355
		{ 
	            // $$ = ast.UnaryExpr{0, $1}
            GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string)) 
	        }
	case 56:
		//line src/gutscript/parser.y:359
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 57:
		//line src/gutscript/parser.y:362
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 58:
		//line src/gutscript/parser.y:367
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 59:
		//line src/gutscript/parser.y:371
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 60:
		//line src/gutscript/parser.y:375
		{
	        GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	    }
	case 61:
		//line src/gutscript/parser.y:380
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.Node) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.Node))
	            GutsVAL.val = params
	        }
	    }
	case 62:
		//line src/gutscript/parser.y:387
		{
	        // create the expr list
        GutsVAL.val = []ast.Node{GutsS[Gutspt-0].val}
	    }
	case 63:
		//line src/gutscript/parser.y:392
		{
	        GutsVAL.val = []ast.Node{}
	    }
	case 64:
		//line src/gutscript/parser.y:398
		{
	        GutsVAL.val = ast.CreateFunctionCall(GutsS[Gutspt-3].val.(string), GutsS[Gutspt-1].val.([]ast.Node))
	    }
	}
	goto Gutsstack /* stack new state and value */
}
