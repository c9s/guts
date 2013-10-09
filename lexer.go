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

	// XXX
	width int
	items chan CoffeeSymType
}

type TokenType int

const eof = -1

func (self *CoffeeSymType) String() string {
	switch self.typ {
	case eof:
		return "EOF"
	}
	return fmt.Sprintf("%q", self.val)
}

/*
func lexText(l *CoffeeLex) stateFn {
    for {
        if strings.HasPrefix(l.input[l.pos:], leftMeta) {
            if l.pos > l.start {
                l.emit(itemText)
            }
            return lexLeftMeta    // Next state.
        }
        if l.next() == eof { break }
    }
    // Correctly reached EOF.
    if l.pos > l.start {
        l.emit(itemText)
    }
    l.emit(itemEOF)  // Useful to make EOF a token.
    return nil       // Stop the run loop.
}
*/

func (l *CoffeeLex) emit(t TokenType) {
	l.items <- CoffeeSymType{
		typ: t,
		val: l.input[l.start:l.pos],
	}
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *CoffeeLex) peek() rune {
	rune := l.next()
	l.backup()
	return rune
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
		l.state = l.state(l)
	}
	close(l.items)
}

func lexStart(l *CoffeeLex) stateFn {
	var c rune = l.next()
	if unicode.IsDigit(c) {
		return lexDigits
	}
	if c == ' ' {
		fmt.Println("run lexSpaces")
		return lexSpaces
	}
	return nil
}

func lexSpaces(l *CoffeeLex) stateFn {
	for c := l.next(); c == ' '; {
	}
	l.backup()
	l.emit(T_SPACE)
	return lexStart
}

func lexDigits(l *CoffeeLex) stateFn {
	return nil
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
