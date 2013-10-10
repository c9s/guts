package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*CoffeeLex) stateFn

type CoffeeLex struct {
	// the line
	input string
	line  int
	start int
	pos   int
	state stateFn

	// current space prefix length (used to calculate indentation level)
	space     int
	lastSpace int

	// XXX
	width int
	items chan *CoffeeSymType
}

type TokenType int

const eof = -1
const lexError = -99

func (self *CoffeeSymType) String() string {
	switch self.typ {
	case T_EOF:
		return "EOF"
	}
	return fmt.Sprintf("[%d] %q", self.typ, self.val)
}

func (l *CoffeeLex) error(msg string) {
	fmt.Println("Syntax Error", msg)
	fmt.Println("Line", l.line)
	fmt.Println("Pos", l.pos)
	fmt.Println("Input", l.input)
}

func (l *CoffeeLex) emit(t TokenType) {
	l.items <- &CoffeeSymType{
		typ:  t,
		val:  l.input[l.start:l.pos],
		pos:  l.start,
		line: l.line,
	}
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *CoffeeLex) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

func (l *CoffeeLex) peekMore(p int) (r rune) {
	var w = 0
	for i := p; i > 0; i-- {
		r = l.next()
		w += l.width
	}
	l.pos -= w
	return r
}

// ignore skips over the pending input before this point.
func (l *CoffeeLex) ignore() {
	l.start = l.pos
}

// accept consumes the next rune
// if it's from the valid set.
// e.g.,
//    l.accept("1234567890")
func (l *CoffeeLex) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *CoffeeLex) backup() {
	l.pos -= l.width
}

// next returns the next rune in the input.
func (l *CoffeeLex) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// get the last rune in the input
func (l *CoffeeLex) last() (r rune) {
	_, w := utf8.DecodeRuneInString(l.input[l.pos-l.width:])
	r, _ = utf8.DecodeRuneInString(l.input[l.pos-l.width-w:])
	return r
}

func (l *CoffeeLex) run() {
	for l.state = lexStart; l.state != nil; {
		fn := l.state(l)
		if fn != nil {
			l.state = fn
		} else {
			break
		}
	}
	l.emit(T_EOF)
}

func (l *CoffeeLex) close() {
	close(l.items)
}

func (l *CoffeeLex) emitIfMatch(str string, typ TokenType) bool {
	if l.consumeIfMatch(str) {
		l.emit(typ)
		return true
	}
	return false
}

func (l *CoffeeLex) consumeIfMatch(str string) bool {
	var c rune
	var width = 0
	for _, sc := range str {
		c = l.next()
		width += l.width
		if sc != c {
			// rollback
			l.pos -= width
			return false
		}
	}
	return true
}

/*
lookahead match method
*/
func (l *CoffeeLex) match(str string) bool {
	var c rune
	var width = 0
	for _, sc := range str {
		c = l.next()
		width += l.width
		if sc != c {
			l.pos -= width
			return false
		}
	}
	l.pos -= width
	return true
}

// set token in lval, return the token type id
func (l *CoffeeLex) Lex(lval *CoffeeSymType) int {
	var c rune = l.next()
	if unicode.IsDigit(c) {
		lval.val = string(c)
		lval.typ = T_DIGIT
		return T_DIGIT
	} else if unicode.IsLower(c) {
		lval.val = string(c)
		lval.typ = T_LETTER
		return T_LETTER
	}
	return int(c)
}

func (l *CoffeeLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
