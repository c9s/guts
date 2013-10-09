
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
	val int
}

const DIGIT = 57346
const LETTER = 57347
const UMINUS = 57348

var CoffeeToknames = []string{
	"DIGIT",
	"LETTER",
	" &&",
	" and",
	" ||",
	" or",
	" |",
	" &",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UMINUS",
}
var CoffeeStatenames = []string{}

const CoffeeEofCode = 1
const CoffeeErrCode = 2
const CoffeeMaxDepth = 200

//line grammer.y:112
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
	20, 22,
	-2, 20,
}

const CoffeeNprod = 28
const CoffeePrivate = 57344

var CoffeeTokenNames []string
var CoffeeStates []string

const CoffeeLast = 73

var CoffeeAct = []int{

	8, 22, 21, 9, 16, 17, 18, 19, 20, 42,
	41, 24, 26, 39, 23, 28, 27, 29, 30, 31,
	32, 33, 34, 35, 36, 22, 21, 38, 16, 17,
	18, 19, 20, 40, 15, 14, 12, 14, 25, 16,
	17, 18, 19, 20, 37, 11, 7, 11, 6, 18,
	19, 20, 10, 5, 10, 22, 21, 4, 16, 17,
	18, 19, 20, 21, 3, 16, 17, 18, 19, 20,
	2, 1, 13,
}
var CoffeePact = []int{

	-1000, -1000, -1000, 31, -1000, -1000, -1000, -1000, 15, -6,
	33, 33, 12, 11, -1000, -1000, 33, 33, 33, 33,
	33, 33, 33, 31, -9, -1000, -1000, -1000, -1000, 34,
	34, -1000, -1000, -1000, 26, 52, 45, -1000, -11, -1000,
	-1000, -13, -1000,
}
var CoffeePgo = []int{

	0, 0, 72, 71, 70, 64, 57, 53, 48, 46,
	3, 44, 33,
}
var CoffeeR1 = []int{

	0, 3, 5, 4, 4, 6, 7, 7, 8, 9,
	9, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 10, 10, 2, 2, 12, 11,
}
var CoffeeR2 = []int{

	0, 1, 0, 3, 0, 1, 1, 1, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	1, 1, 1, 2, 1, 2, 2, 2,
}
var CoffeeChk = []int{

	-1000, -3, -4, -5, -6, -7, -8, -9, -1, -10,
	21, 14, 5, -2, 4, 19, 13, 14, 15, 16,
	17, 11, 10, 20, -1, 5, -1, 4, 4, -1,
	-1, -1, -1, -1, -1, -1, -1, -11, -10, 22,
	-12, 21, 22,
}
var CoffeeDef = []int{

	4, -2, -2, 0, 3, 5, 6, 7, 0, 0,
	0, 0, -2, 21, 24, 8, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 20, 19, 23, 25, 12,
	13, 14, 15, 16, 17, 18, 9, 10, 0, 11,
	27, 0, 26,
}
var CoffeeTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 17, 11, 3,
	21, 22, 15, 13, 3, 14, 3, 16, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 19,
	3, 20, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 12, 3, 3, 7, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 9, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 10,
}
var CoffeeTok2 = []int{

	2, 3, 4, 5, 18,
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
		//line grammer.y:38
		{ }
	case 2:
		//line grammer.y:41
		{ }
	case 3:
		//line grammer.y:41
		{ }
	case 5:
		//line grammer.y:46
		{ }
	case 6:
		//line grammer.y:50
		{ }
	case 7:
		//line grammer.y:51
		{ }
	case 8:
		//line grammer.y:55
		{ }
	case 9:
		//line grammer.y:59
		{  }
	case 10:
		//line grammer.y:60
		{  }
	case 11:
		//line grammer.y:64
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-1].val }
	case 12:
		//line grammer.y:66
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val + CoffeeS[Coffeept-0].val }
	case 13:
		//line grammer.y:68
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val - CoffeeS[Coffeept-0].val }
	case 14:
		//line grammer.y:70
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val * CoffeeS[Coffeept-0].val }
	case 15:
		//line grammer.y:72
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val / CoffeeS[Coffeept-0].val }
	case 16:
		//line grammer.y:74
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val % CoffeeS[Coffeept-0].val }
	case 17:
		//line grammer.y:76
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val & CoffeeS[Coffeept-0].val }
	case 18:
		//line grammer.y:78
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-2].val | CoffeeS[Coffeept-0].val }
	case 19:
		//line grammer.y:80
		{ CoffeeVAL.val  = -CoffeeS[Coffeept-0].val  }
	case 20:
		//line grammer.y:82
		{ CoffeeVAL.val  = regs[CoffeeS[Coffeept-0].val] }
	case 21:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 24:
		//line grammer.y:92
		{
				CoffeeVAL.val = CoffeeS[Coffeept-0].val;
				if CoffeeS[Coffeept-0].val==0 {
					base = 8
				} else {
					base = 10
				}
			}
	case 25:
		//line grammer.y:101
		{ CoffeeVAL.val = base * CoffeeS[Coffeept-1].val + CoffeeS[Coffeept-0].val }
	case 26:
		//line grammer.y:105
		{ }
	case 27:
		//line grammer.y:109
		{ }
	}
	goto Coffeestack /* stack new state and value */
}
