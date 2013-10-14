
//line src/gutscript/parser.y:2
package gutscript
import __yyfmt__ "fmt"
//line src/gutscript/parser.y:2
		
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


//line src/gutscript/parser.y:28
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

//line src/gutscript/parser.y:348
      /*  start  of  programs  */

//line yacctab:1
var GutsExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const GutsNprod = 46
const GutsPrivate = 57344

var GutsTokenNames []string
var GutsStates []string

const GutsLast = 171

var GutsAct = []int{

	66, 62, 59, 28, 5, 27, 22, 23, 24, 25,
	26, 22, 23, 24, 25, 26, 11, 56, 29, 30,
	29, 30, 63, 29, 30, 69, 4, 70, 41, 28,
	40, 27, 22, 23, 24, 25, 26, 53, 37, 31,
	32, 67, 11, 67, 29, 30, 28, 60, 27, 22,
	23, 24, 25, 26, 24, 25, 26, 61, 71, 1,
	2, 29, 30, 55, 29, 30, 9, 7, 68, 28,
	72, 27, 22, 23, 24, 25, 26, 27, 22, 23,
	24, 25, 26, 36, 29, 30, 17, 20, 19, 64,
	29, 30, 65, 11, 54, 58, 17, 20, 19, 18,
	35, 20, 19, 11, 8, 16, 15, 34, 14, 18,
	14, 3, 0, 0, 21, 0, 10, 0, 0, 14,
	0, 0, 0, 0, 0, 0, 10, 13, 0, 6,
	0, 0, 0, 0, 0, 0, 12, 13, 0, 0,
	33, 13, 38, 39, 14, 0, 12, 0, 42, 0,
	12, 0, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 0, 0, 0, 0,
	57,
}
var GutsPact = []int{

	91, -1000, 91, -1000, -1000, -1000, 27, -1000, -1000, 20,
	95, 91, 95, 95, 15, -1000, -1000, -3, 95, -1000,
	-1000, -1000, 95, 95, 95, 95, 95, 95, 95, 95,
	95, 95, 30, 27, -1000, -1000, 81, -1000, -39, -37,
	95, -53, 4, 7, 7, -37, -37, -37, -34, 33,
	27, 27, 4, -1000, -1000, -1000, -1000, -13, -10, 36,
	-1000, -1000, -1000, 30, -29, -1000, -1000, 53, -1000, 38,
	-1000, -1000, -1000,
}
var GutsPgo = []int{

	0, 129, 107, 106, 105, 26, 104, 95, 89, 0,
	83, 67, 66, 4, 111, 60, 59, 59, 59,
}
var GutsR1 = []int{

	0, 16, 15, 15, 14, 13, 10, 10, 5, 5,
	5, 5, 5, 5, 12, 12, 12, 11, 11, 9,
	9, 8, 8, 7, 7, 7, 6, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 4, 3, 17, 18,
}
var GutsR2 = []int{

	0, 1, 2, 1, 1, 3, 2, 1, 1, 1,
	1, 1, 1, 2, 3, 4, 3, 3, 4, 2,
	1, 3, 1, 3, 2, 0, 5, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
	1, 1, 1, 1, 2, 2,
}
var GutsChk = []int{

	-1000, -16, -15, -14, -5, -13, -1, -11, -6, -12,
	35, 12, 55, 46, -2, -3, -4, 5, 18, 7,
	6, -14, 45, 46, 47, 48, 49, 44, 42, 57,
	58, 19, 20, -1, -2, 5, -10, -5, -1, -1,
	15, 31, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -13, 13, -5, 56, -1, -7, 55,
	-13, -13, 14, 32, -8, 56, -9, 5, -13, 54,
	56, 5, -9,
}
var GutsDef = []int{

	0, -2, 1, 3, 4, 8, 9, 10, 11, 12,
	0, 0, 0, 0, 39, 40, 41, 27, 0, 43,
	42, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 39, 27, 0, 7, 0, 38,
	0, 25, 0, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 0, 16, 5, 6, 28, 17, 0, 0,
	14, 15, 18, 0, 0, 24, 22, 20, 26, 0,
	23, 19, 21,
}
var GutsTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 3, 3, 3, 49, 44, 3,
	55, 56, 47, 45, 54, 46, 3, 48, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	58, 3, 57, 3, 3, 3, 3, 3, 3, 3,
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
		//line src/gutscript/parser.y:129
		{
	        debug("end compilation", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 2:
		//line src/gutscript/parser.y:136
		{ 
	        if stmts, ok := GutsS[Gutspt-1].val.(*ast.StatementNodeList) ; ok {
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = GutsS[Gutspt-1].val
	        }
	        debug("top_statement_list top_statement", GutsS[Gutspt-1].val, GutsS[Gutspt-0].val)
	    }
	case 3:
		//line src/gutscript/parser.y:143
		{
	        debug("top_statement", GutsS[Gutspt-0].val)
	        stmts := ast.CreateStatementList()
	        stmts.Append(GutsS[Gutspt-0].val)
	        GutsVAL.val = stmts
	    }
	case 4:
		//line src/gutscript/parser.y:152
		{
	        debug("statement", GutsS[Gutspt-0].val)
	        GutsVAL.val = GutsS[Gutspt-0].val
	    }
	case 5:
		//line src/gutscript/parser.y:160
		{ GutsVAL.val = GutsS[Gutspt-1].val }
	case 6:
		//line src/gutscript/parser.y:164
		{
	            if stmts, ok := GutsS[Gutspt-1].val.(*ast.StatementNodeList) ; ok {
	                stmts.Append(GutsS[Gutspt-0].val)
	                GutsVAL.val = GutsS[Gutspt-1].val
	            }
	      }
	case 7:
		//line src/gutscript/parser.y:171
		{ 
	            stmts := ast.CreateStatementList()
	            stmts.Append(GutsS[Gutspt-0].val)
	            GutsVAL.val = stmts
	      }
	case 8:
		//line src/gutscript/parser.y:179
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 9:
		//line src/gutscript/parser.y:180
		{ GutsVAL.val = ast.CreateExprStatement(GutsS[Gutspt-0].val) }
	case 10:
		//line src/gutscript/parser.y:181
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 11:
		//line src/gutscript/parser.y:182
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 12:
		//line src/gutscript/parser.y:183
		{ GutsVAL.val = GutsS[Gutspt-0].val }
	case 13:
		//line src/gutscript/parser.y:184
		{ GutsVAL.val = ast.CreateReturnStatement(GutsS[Gutspt-0].val) }
	case 14:
		//line src/gutscript/parser.y:189
		{
	            GutsVAL.val = ast.CreateIfStatement(GutsS[Gutspt-1].val.(ast.Expr), GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	        }
	case 15:
		//line src/gutscript/parser.y:194
		{
	            GutsS[Gutspt-3].val.(*ast.IfStatement).AddElseIf(GutsS[Gutspt-1].val.(ast.Expr),GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	            GutsVAL.val = GutsS[Gutspt-3].val
	        }
	case 16:
		//line src/gutscript/parser.y:200
		{
	            GutsS[Gutspt-2].val.(*ast.IfStatement).SetElse(GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	            GutsVAL.val = GutsS[Gutspt-2].val
	        }
	case 17:
		//line src/gutscript/parser.y:208
		{
	            debug("assignment_statement", GutsS[Gutspt-2].val , "=" , GutsS[Gutspt-0].val)
	            GutsVAL.val = ast.CreateAssignStatement(GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 18:
		//line src/gutscript/parser.y:213
		{
	            debug("assignment_statement", GutsS[Gutspt-3].val , "=" , GutsS[Gutspt-1].val)
	            GutsVAL.val = ast.CreateAssignStatement(GutsS[Gutspt-3].val, GutsS[Gutspt-1].val)
	        }
	case 19:
		//line src/gutscript/parser.y:227
		{
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), GutsS[Gutspt-1].val.(string), nil}
	    }
	case 20:
		//line src/gutscript/parser.y:231
		{ 
	        GutsVAL.val = ast.FunctionParam{GutsS[Gutspt-0].val.(string), "", nil}
	    }
	case 21:
		//line src/gutscript/parser.y:237
		{
	        if params, ok := GutsS[Gutspt-2].val.([]ast.FunctionParam) ; ok {
	            params = append(params, GutsS[Gutspt-0].val.(ast.FunctionParam))
	            GutsVAL.val = params
	        }
	      }
	case 22:
		//line src/gutscript/parser.y:243
		{
	            GutsVAL.val = []ast.FunctionParam{GutsS[Gutspt-0].val.(ast.FunctionParam)}
	        }
	case 23:
		//line src/gutscript/parser.y:249
		{
	        GutsVAL.val = GutsS[Gutspt-1].val
	    }
	case 24:
		//line src/gutscript/parser.y:252
		{
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 25:
		//line src/gutscript/parser.y:255
		{ 
	        GutsVAL.val = []ast.FunctionParam{}
	    }
	case 26:
		//line src/gutscript/parser.y:261
		{ 
	        GutsVAL.val = ast.CreateFunction(GutsS[Gutspt-4].val.(string), GutsS[Gutspt-2].val.([]ast.FunctionParam), GutsS[Gutspt-0].val.(*ast.StatementNodeList))
	    }
	case 27:
		//line src/gutscript/parser.y:267
		{ 
	        GutsVAL.val = ast.CreateVariable(GutsS[Gutspt-0].val.(string))
	    }
	case 28:
		//line src/gutscript/parser.y:272
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
		//line src/gutscript/parser.y:282
		{ 
	            GutsVAL.val = ast.CreateExpr('+', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 30:
		//line src/gutscript/parser.y:286
		{ 
	            GutsVAL.val = ast.CreateExpr('-', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 31:
		//line src/gutscript/parser.y:290
		{ 
	            GutsVAL.val = ast.CreateExpr('*', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 32:
		//line src/gutscript/parser.y:294
		{ 
	            GutsVAL.val = ast.CreateExpr('/', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 33:
		//line src/gutscript/parser.y:298
		{ 
	            GutsVAL.val = ast.CreateExpr('%', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 34:
		//line src/gutscript/parser.y:302
		{
	            GutsVAL.val = ast.CreateExpr('&', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 35:
		//line src/gutscript/parser.y:306
		{
	            GutsVAL.val = ast.CreateExpr('|', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 36:
		//line src/gutscript/parser.y:310
		{
	            GutsVAL.val = ast.CreateExpr('>', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 37:
		//line src/gutscript/parser.y:314
		{
	            GutsVAL.val = ast.CreateExpr('<', GutsS[Gutspt-2].val, GutsS[Gutspt-0].val)
	        }
	case 38:
		//line src/gutscript/parser.y:318
		{ 
	            GutsVAL.val = ast.UnaryExpr{'-', GutsS[Gutspt-0].val}
	        }
	case 39:
		//line src/gutscript/parser.y:321
		{ 
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 40:
		//line src/gutscript/parser.y:324
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 41:
		//line src/gutscript/parser.y:327
		{
	            GutsVAL.val = ast.UnaryExpr{0, GutsS[Gutspt-0].val}
	        }
	case 42:
		//line src/gutscript/parser.y:332
		{
	        GutsVAL.val = ast.CreateFloatingNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 43:
		//line src/gutscript/parser.y:336
		{
	        GutsVAL.val = ast.CreateNumber(GutsS[Gutspt-0].val.(string))
	    }
	case 44:
		//line src/gutscript/parser.y:341
		{ }
	case 45:
		//line src/gutscript/parser.y:345
		{ }
	}
	goto Gutsstack /* stack new state and value */
}
