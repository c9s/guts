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

func (l *CoffeeLex) emit(t TokenType) {
	l.items <- &CoffeeSymType{
		typ: t,
		val: l.input[l.start:l.pos],
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

func lexStartLine(l *CoffeeLex) stateFn {
	var c rune = l.peek()
	if c == ' ' || c == '\t' {
		return lexIndentSpaces
	}
	// default start state.
	return lexStart
}

func lexStart(l *CoffeeLex) stateFn {
	var c rune = l.peek()
	if unicode.IsDigit(c) {
		return lexNumber
	} else if c == ' ' || c == '\t' {
		// return lexSpaces
		return lexIgnoreSpaces
	} else if c == '\n' || c == '\r' {
		l.next()
		l.emit(T_NEWLINE)
		l.lastSpace = l.space
		l.space = 0
		return lexStartLine
	} else if c == '=' && l.peekMore(2) != '=' {
		l.next()
		l.emit(T_ASSIGN)
		return lexStart
	} else if l.consumeIfMatch("//") {
		return lexOnelineComment
	} else if l.consumeIfMatch("/*") {
		return lexComment
	} else if l.consumeIfMatch("if ") {
		l.emit(T_IF)
		return lexStart
	} else if l.consumeIfMatch("if else ") {
		l.emit(T_ELSEIF)
		return lexStart
	} else if unicode.IsLetter(c) {
		return lexIdentifier
	} else if c == eof {
		l.emit(T_EOF)
		return nil
	}
	return nil
}

func lexComment(l *CoffeeLex) stateFn {
	var c rune
	for {
		c = l.next()
		if c == eof {
			break
		}
		if c == '*' && l.peek() == '/' {
			break
		}
	}
	l.next()
	l.emit(T_COMMENT)
	return lexStart
}

func lexOnelineComment(l *CoffeeLex) stateFn {
	var c rune
	for {
		c = l.next()
		if c == '\n' || c == eof {
			break
		}
	}
	l.backup()
	l.emit(T_ONELINE_COMMENT)
	return lexStartLine
}

func lexIdentifier(l *CoffeeLex) stateFn {
	for {
		c := l.next()
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
		} else {
			break
		}
	}
	l.backup()
	l.emit(T_IDENTIFIER)
	return lexStart
}

func lexIgnoreSpaces(l *CoffeeLex) stateFn {
	var c rune
	for {
		c = l.next()
		if c != ' ' || c == eof {
			break
		}
	}
	l.backup()
	l.ignore()
	return lexStart
}

func lexIndentSpaces(l *CoffeeLex) stateFn {
	l.space = 0
	for {
		c := l.next()
		if c == ' ' || c == '\t' {
			l.space++
		} else {
			break
		}
	}
	l.backup()
	l.emit(T_SPACE)
	return lexStart
}

func lexSpaces(l *CoffeeLex) stateFn {
	for c := l.next(); c == ' ' || c == '\t'; {
	}
	l.backup()
	l.emit(T_SPACE)
	return lexStart
}

func lexFloating(l *CoffeeLex) stateFn {
	for {
		c := l.next()
		if !unicode.IsDigit(c) {
			break
		}
	}
	l.backup()
	l.emit(T_FLOATING)
	return lexStart
}

func lexNumber(l *CoffeeLex) stateFn {
	var c rune
	for c = l.next(); true; c = l.next() {
		if unicode.IsDigit(c) {
			continue
		} else if c == '.' && unicode.IsDigit(l.peek()) {
			return lexFloating
		} else {
			break
		}
	}
	l.backup()
	l.emit(T_NUMBER)
	return lexStart
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
