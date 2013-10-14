
//line src/gutscript/parser.y:2
package gutscript
import __yyfmt__ "fmt"
//line src/gutscript/parser.y:2
		
// vim:et:sw=4:ai:si:ts=4:sts=4:

import "fmt"
import "gutscript/ast"

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
const T_FUNCTION_GLYPH = 57374
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
	"T_FUNCTION_GLYPH",
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
var GutsStatenames = []string{}

const GutsEofCode = 1
const GutsErrCode = 2
const GutsMaxDepth = 200

//line src/gutscript/parser.y:312
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 41
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 178

var GutsAct = []int{

	5, 29, 30, 28, 62, 27, 22, 23, 24, 25,
	26, 24, 25, 26, 64, 63, 56, 29, 30, 59,
	29, 30, 11, 41, 31, 32, 40, 11, 34, 14,
	1, 14, 28, 53, 27, 22, 23, 24, 25, 26,
	14, 3, 2, 60, 21, 9, 29, 30, 7, 36,
	4, 58, 28, 61, 27, 22, 23, 24, 25, 26,
	8, 16, 37, 15, 65, 14, 29, 30, 28, 0,
	27, 22, 23, 24, 25, 26, 27, 22, 23, 24,
	25, 26, 29, 30, 35, 20, 19, 55, 29, 30,
	22, 23, 24, 25, 26, 0, 17, 20, 19, 0,
	0, 29, 30, 11, 54, 17, 20, 19, 0, 18,
	0, 0, 11, 0, 0, 0, 0, 0, 18, 0,
	0, 0, 0, 0, 0, 13, 10, 0, 0, 0,
	0, 0, 0, 12, 0, 10, 6, 13, 0, 0,
	0, 0, 0, 0, 0, 12, 13, 33, 0, 38,
	39, 0, 0, 0, 12, 42, 0, 0, 0, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 0, 0, 0, 0, 57,
}
var GutsPact = []int{

	100, -1000, 100, -1000, -1000, -1000, 26, -1000, -1000, 5,
	79, 100, 79, 79, 11, -1000, -1000, -8, 79, -1000,
	-1000, -1000, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 15, 26, -1000, -1000, 91, -1000, -39, -55,
	79, -35, 10, -36, -36, -55, -55, -55, 45, 32,
	26, 26, 10, -1000, -1000, -1000, -1000, -10, -17, -41,
	-1000, -1000, -1000, 15, -1000, -1000,
}
var GutsPgo = []int{

	0, 136, 28, 63, 61, 50, 60, 51, 49, 48,
	45, 0, 41, 42, 30, 30, 30,
}
var GutsR1 = []int{

	0, 14, 13, 13, 12, 11, 8, 8, 5, 5,
	5, 5, 5, 5, 10, 10, 10, 9, 9, 7,
	7, 6, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 3, 15,
	16,
}
var GutsR2 = []int{

	0, 1, 2, 1, 1, 3, 2, 1, 1, 1,
	1, 1, 1, 2, 3, 4, 3, 3, 4, 2,
	0, 5, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 1, 1, 1, 1, 1, 2,
	2,
}
var GutsChk = []int{

	-1000, -14, -13, -12, -5, -11, -1, -9, -6, -10,
	35, 12, 54, 46, -2, -3, -4, 5, 18, 7,
	6, -12, 45, 46, 47, 48, 49, 44, 42, 56,
	57, 19, 20, -1, -2, 5, -8, -5, -1, -1,
	15, 31, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -11, 13, -5, 55, -1, -7, 54,
	-11, -11, 14, 32, 55, -11,
}
var GutsDef = []int{

	0, -2, 1, 3, 4, 8, 9, 10, 11, 12,
	0, 0, 0, 0, 34, 35, 36, 22, 0, 38,
	37, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 34, 22, 0, 7, 0, 33,
	0, 20, 0, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 0, 16, 5, 6, 23, 17, 0, 0,
	14, 15, 18, 0, 19, 21,
}
var GutsTok1 = []int{

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
var GutsTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	51, 52, 53,
}
var GutsTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var GutsDebug = 0

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
		//line src/gutscript/parser.y:125
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:132
		{ 
	        if stmts, ok := GutsS[Gutspt-1].val.(*ast.StatementNodeList) ; ok {
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = GutsS[Gutspt-1].val
	        }
	        debug("top_statement_list top_statement", GutsS[Gutspt-1].val, GutsS[Gutspt-0].val)
	    }
	case 3:
		//line src/gutscript/parser.y:139
		{
	        debug("top_statement", GutsS[Gutspt-0].val)
	        stmts := ast.CreateStatementList()
	        stmts.Append(GutsS[Gutspt-0].val)
	        GutsVAL.val = stmts
	    }
	case 4:
		//line src/gutscript/parser.y:148
		{
	        debug("statement", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 5:
		//line src/gutscript/parser.y:156
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 6:
		//line src/gutscript/parser.y:160
		{
	            if stmts, ok := GutsS[Gutspt-1].val.(*ast.StatementNodeList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-1].val
	            }
	      }
	case 7:
		//line src/gutscript/parser.y:167
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 8:
		//line src/gutscript/parser.y:175
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 9:
		//line src/gutscript/parser.y:176
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 10:
		//line src/gutscript/parser.y:177
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:178
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:179
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:180
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 14:
		//line src/gutscript/parser.y:185
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	        }
	case 15:
		//line src/gutscript/parser.y:190
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 16:
		//line src/gutscript/parser.y:196
		{
	            GutsS[Gutspt-2].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	            GutsVAL.val = GutsS[Gutspt-2].val
	        }
	case 17:
		//line src/gutscript/parser.y:204
		{
	            debug("assignment_statement", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 18:
		//line src/gutscript/parser.y:209
		{
	            debug("assignment_statement", GutsS[Gutspt-3].val , "=" , GutsS[Gutspt-1].val)
	            GutsVAL.val = ast.CreateAssignStatement(GutsS[Gutspt-3].val, GutsS[Gutspt-1].val)
	        }
	case 19:
		//line src/gutscript/parser.y:216
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 20:
		//line src/gutscript/parser.y:219
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 21:
		//line src/gutscript/parser.y:225
		{ 
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	    }
	case 22:
		//line src/gutscript/parser.y:231
		{ 
	        GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string))
	    }
	case 23:
		//line src/gutscript/parser.y:236
		{
	            if node, ok := GutsS[Gutspt-1].val.(ast.Expr) ; ok {
	                node.Parenthesis = true
	                GutsVAL.val = node
	            } else {
	                panic(" type casting to ast.Expr failed.")
	            }
	            // $$ = $2
        }
	case 24:
		//line src/gutscript/parser.y:246
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 25:
		//line src/gutscript/parser.y:250
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 26:
		//line src/gutscript/parser.y:254
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 27:
		//line src/gutscript/parser.y:258
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 28:
		//line src/gutscript/parser.y:262
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 29:
		//line src/gutscript/parser.y:266
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 30:
		//line src/gutscript/parser.y:270
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 31:
		//line src/gutscript/parser.y:274
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 32:
		//line src/gutscript/parser.y:278
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 33:
		//line src/gutscript/parser.y:282
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 34:
		//line src/gutscript/parser.y:285
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 35:
		//line src/gutscript/parser.y:288
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 36:
		//line src/gutscript/parser.y:291
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 37:
		//line src/gutscript/parser.y:296
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 38:
		//line src/gutscript/parser.y:300
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 39:
		//line src/gutscript/parser.y:305
		{ }
	case 40:
		//line src/gutscript/parser.y:309
		{ }
	}
	goto Gutsstack /* stack new state and value */
}
