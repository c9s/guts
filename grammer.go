
//line grammer.y:2

package main
import __yyfmt__ "fmt"
//line grammer.y:3
		
import _ "fmt"

var regs = make([]int, 26)
var base int


//line grammer.y:14
type CoffeeSymType struct{
	yys int
    typ TokenType
	val string
    line int
    pos  int
}

const T_DIGIT = 57346
const T_LETTER = 57347
const T_DOT = 57348
const T_IDENTIFIER = 57349
const T_EOF = 57350
const T_FLOATING = 57351
const T_NUMBER = 57352
const T_ONELINE_COMMENT = 57353
const T_COMMENT = 57354
const T_NEWLINE = 57355
const T_ASSIGN = 57356
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
const T_STRING = 57373
const T_CONST = 57374
const T_RETURN = 57375
const T_BREAK = 57376
const T_CONTINUE = 57377
const T_THROW = 57378
const T_NAMESPACE = 57379
const T_OBJECT_OPERATOR = 57380
const T_BOOLEAN_OR = 57381
const T_BOOLEAN_AND = 57382
const UMINUS = 57383

var CoffeeToknames = []string{
	"T_DIGIT",
	"T_LETTER",
	"T_DOT",
	"T_IDENTIFIER",
	"T_EOF",
	"T_FLOATING",
	"T_NUMBER",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
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
	" class (T_CLASS)",
	"T_IS",
	" is (T_IS)",
	"T_DOES",
	" does (T_DOES)",
	"T_FUNCTION_PROTOTYPE",
	" :: (T_FUNCTION_PROTOTYPE)",
	"T_STRING",
	"T_CONST",
	" const (T_CONST)",
	"T_RETURN",
	" return (T_RETURN)",
	"T_BREAK",
	" break (T_BREAK)",
	"T_CONTINUE",
	" continue (T_CONTINUE)",
	"T_THROW",
	" throw (T_THROW)",
	"T_NAMESPACE",
	" namespace (T_NAMESPACE)",
	"T_OBJECT_OPERATOR",
	" . (T_OBJECT_OPERATOR)",
	" and",
	" or",
	" |",
	" ^",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	" !",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"UMINUS",
}
var CoffeeStatenames = []string{}

const CoffeeEofCode = 1
const CoffeeErrCode = 2
const CoffeeMaxDepth = 200

//line grammer.y:202
      /*  start  of  programs  */

//line yacctab:1
var CoffeeExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 2,
	-1, 12,
	65, 25,
	66, 25,
	-2, 23,
}

const CoffeeNprod = 31
const CoffeePrivate = 57344

var CoffeeTokenNames []string
var CoffeeStates []string

const CoffeeLast = 82

var CoffeeAct = []int{

	22, 44, 21, 16, 17, 18, 19, 20, 14, 12,
	14, 25, 43, 41, 22, 39, 21, 16, 17, 18,
	19, 20, 23, 28, 27, 22, 40, 21, 16, 17,
	18, 19, 20, 18, 19, 20, 42, 15, 21, 16,
	17, 18, 19, 20, 16, 17, 18, 19, 20, 37,
	7, 6, 5, 4, 3, 2, 8, 9, 1, 13,
	11, 0, 11, 0, 0, 0, 0, 24, 26, 0,
	10, 0, 10, 29, 30, 31, 32, 33, 34, 35,
	36, 38,
}
var CoffeePact = []int{

	-1000, -1000, -1000, 4, -1000, -1000, -1000, -1000, -27, -43,
	6, 6, 20, 19, -1000, -1000, 6, 6, 6, 6,
	6, 6, 6, 4, -52, -1000, -1000, -1000, -1000, -24,
	-24, -1000, -1000, -1000, -11, -16, -38, -51, -54, -1000,
	-1000, -1000, -1000, -66, -1000,
}
var CoffeePgo = []int{

	0, 56, 59, 58, 55, 54, 53, 52, 51, 50,
	57, 49, 49, 49, 49, 36,
}
var CoffeeR1 = []int{

	0, 3, 5, 4, 4, 6, 7, 7, 8, 9,
	9, 12, 13, 14, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 10, 10, 2, 2, 15,
	11,
}
var CoffeeR2 = []int{

	0, 1, 0, 3, 0, 1, 1, 1, 2, 4,
	4, 2, 5, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1, 1, 2, 1, 2, 2,
	2,
}
var CoffeeChk = []int{

	-1000, -3, -4, -5, -6, -7, -8, -9, -1, -10,
	66, 56, 5, -2, 4, 64, 55, 56, 57, 58,
	59, 54, 52, 65, -1, 5, -1, 4, 4, -1,
	-1, -1, -1, -1, -1, -1, -1, -11, -10, 67,
	64, 64, -15, 66, 67,
}
var CoffeeDef = []int{

	4, -2, -2, 0, 3, 5, 6, 7, 0, 0,
	0, 0, -2, 24, 27, 8, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 23, 22, 26, 28, 15,
	16, 17, 18, 19, 20, 21, 0, 0, 0, 14,
	9, 10, 30, 0, 29,
}
var CoffeeTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 60, 3, 3, 3, 59, 54, 3,
	66, 67, 57, 55, 3, 68, 49, 58, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 34, 64,
	3, 65, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 53, 3, 3, 50, 41, 43,
	32, 3, 3, 3, 3, 30, 3, 3, 3, 3,
	47, 51, 3, 3, 39, 3, 45, 3, 3, 3,
	3, 3, 3, 3, 52,
}
var CoffeeTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 29, 31, 33, 35,
	36, 38, 40, 42, 44, 46, 48, 61, 62, 63,
}
var CoffeeTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var CoffeeDebug = 0

type CoffeeLexer interface {
	Lex(lval *CoffeeSymType) int
	Error(s string)
}

const CoffeeFlag = -1000

func CoffeeTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(CoffeeToknames) {
		if CoffeeToknames[c-4] != "" {
			return CoffeeToknames[c-4]
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

func CoffeeParse(Coffeelex CoffeeLexer) int {
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
		//line grammer.y:96
		{ }
	case 2:
		//line grammer.y:99
		{ }
	case 3:
		//line grammer.y:99
		{ }
	case 5:
		//line grammer.y:104
		{ }
	case 6:
		//line grammer.y:108
		{ }
	case 7:
		//line grammer.y:109
		{ }
	case 8:
		//line grammer.y:113
		{ }
	case 9:
		//line grammer.y:117
		{  }
	case 10:
		//line grammer.y:118
		{  }
	case 14:
		//line grammer.y:130
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-1].val }
	case 15:
		//line grammer.y:132
		{ 
	            // $$  =  $1 + $3 
        }
	case 16:
		//line grammer.y:136
		{ 
	            // $$  =  $1 - $3 
        }
	case 17:
		//line grammer.y:140
		{ 
	            // $$  =  $1 * $3 
        }
	case 18:
		//line grammer.y:144
		{ 
	            // $$  =  $1 / $3 
        }
	case 19:
		//line grammer.y:148
		{ 
	            // $$  =  $1 % $3 
        }
	case 20:
		//line grammer.y:152
		{ 
	            // $$  =  $1 & $3 
        }
	case 21:
		//line grammer.y:156
		{ 
	            // $$  =  $1 | $3 
        }
	case 22:
		//line grammer.y:160
		{ 
	            // $$  = -$2  
        }
	case 23:
		//line grammer.y:164
		{ 
	            // $$  = regs[$1] 
        }
	case 24:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 27:
		//line grammer.y:179
		{
				CoffeeVAL.val = CoffeeS[Coffeept-0].val;
	            /*
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
            */
		}
	case 28:
		//line grammer.y:189
		{ 
	            // $$ = base * $1 + $2 
        }
	case 29:
		//line grammer.y:195
		{ }
	case 30:
		//line grammer.y:199
		{ }
	}
	goto Coffeestack /* stack new state and value */
}
