package gutscript

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type stateFn func(*GutsLex) stateFn

type GutsLex struct {
	// the line
	Input       string
	Line        int
	Start       int
	Pos         int
	IndentLevel int
	state       stateFn

	// rollback position
	rollbackPos int

	// current space prefix length (used to calculate indentation level)
	space     int
	lastSpace int

	lastTokenType TokenType

	// XXX
	width int
	Items chan *GutsSymType
}

type TokenType int

const eof = -1
const lexError = -99

// remember rollback position
func (l *GutsLex) remember() int {
	l.rollbackPos = l.Pos
	return l.Pos
}

func (l *GutsLex) rollback() {
	l.Pos = l.rollbackPos
}

func (l *GutsLex) error(msg string) {
	fmt.Println("Syntax Error", msg)
	fmt.Println("Line", l.Line)
	fmt.Println("Pos", l.Pos)
	fmt.Println("Input", l.Input)
}

func (l *GutsLex) emit(t TokenType) {

	l.lastTokenType = t
	l.Items <- &GutsSymType{
		typ:  t,
		val:  l.Input[l.Start:l.Pos],
		pos:  l.Start,
		line: l.Line,
	}
	l.Start = l.Pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *GutsLex) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

func (l *GutsLex) peekMore(p int) (r rune) {
	var w = 0
	for i := p; i > 0; i-- {
		r = l.next()
		w += l.width
	}
	l.Pos -= w
	return r
}

// ignore skips over the pending input before this point.
func (l *GutsLex) ignore() {
	l.Start = l.Pos
}

// accept consumes the next rune
// if it's from the valid set.
// e.g.,
//    l.accept("1234567890")
func (l *GutsLex) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *GutsLex) backup() {
	l.Pos -= l.width
}

// next returns the next rune in the input.
func (l *GutsLex) next() (r rune) {
	if l.Pos >= len(l.Input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Pos += l.width
	return r
}

// get the last rune in the input
func (l *GutsLex) last() (r rune) {
	_, w := utf8.DecodeRuneInString(l.Input[l.Pos-l.width:])
	r, _ = utf8.DecodeRuneInString(l.Input[l.Pos-l.width-w:])
	return r
}

func (l *GutsLex) run() {
	for l.state = lexStart; l.state != nil; {
		fn := l.state(l)
		if fn != nil {
			l.state = fn
		} else {
			break
		}
	}
	l.Items <- nil
}

func (l *GutsLex) close() {
	close(l.Items)
}

func (l *GutsLex) emitIfMatch(str string, typ TokenType) bool {
	if l.consumeIfMatch(str) {
		l.emit(typ)
		return true
	}
	return false
}

func (l *GutsLex) consumeIfMatch(str string) bool {
	var c rune
	var width = 0
	for _, sc := range str {
		c = l.next()
		width += l.width
		if sc != c {
			l.Pos -= width
			return false
		}
	}
	return true
}

/*
lookahead match method
*/
func (l *GutsLex) match(str string) bool {
	var c rune
	var width = 0
	for _, sc := range str {
		c = l.next()
		width += l.width
		if sc != c {
			l.Pos -= width
			return false
		}
	}
	l.Pos -= width
	return true
}

// set token in lval, return the token type id
func (l *GutsLex) Lex(lval *GutsSymType) int {
	var item *GutsSymType
	item = <-l.Items
	if item != nil {
		*lval = *item
		// fmt.Println("Lex", item.String())
		// fmt.Printf("%s %s", GutsTokname(int(item.typ)), item.val)
		return int(item.typ)
	}
	return 0
}

func GetTokenName(typ int) string {
	var c int = typ
	if typ <= 0 {
		c = GutsTok1[0]
		goto out2
	}
	if typ < len(GutsTok1) {
		c = GutsTok1[typ]
		goto out2
	}
	if c >= GutsPrivate {
		if c < GutsPrivate+len(GutsTok2) {
			c = GutsTok2[c-GutsPrivate]
			goto out2
		}
	}
	for i := 0; i < len(GutsTok3); i += 2 {
		c = GutsTok3[i+0]
		if c == typ {
			c = GutsTok3[i+1]
			goto out2
		}
	}
out2:
	if c >= 4 && c-4 < len(GutsToknames) {
		if GutsToknames[c-4] != "" {
			return GutsToknames[c-4]
		}
	}
	return fmt.Sprintf("'%c'", typ)
}

func (l *GutsLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func dumpLexItems(items chan *GutsSymType) {
	for {
		item := <-items
		if item == nil {
			break
		}
		fmt.Printf("Got token %s: '%s'\n", GetTokenName(int(item.typ)), item.val)
		if item.typ == T_EOF || item.typ == eof {
			break
		}
	}
}
