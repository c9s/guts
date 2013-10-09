package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type CoffeeLex struct {
	// the line
	input string
	start int
	pos   int

	// XXX
	width int
	items chan LexItem
}

type TokenType int

const eof = -1

type LexItem struct {
	Val string
	Typ TokenType
}

func (self *LexItem) String() string {
	switch self.Typ {
	case eof:
		return "EOF"
	}
	return fmt.Sprintf("%q", self.Val)
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
func (l *CoffeeLex) emit(t TokenType) {
    l.items <- item{t, l.input[l.start:l.pos]}
    l.start = l.pos
}
func (l *CoffeeLex) run() {
	for state := l.lexText; state != nil {
		state = state(l)
	}
	close(l.items)
}
*/

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
	r, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// set token in lval, return the token type id
func (l *CoffeeLex) Lex(lval *CoffeeSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.input) {
			return 0
		}
		c = rune(l.input[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c - '0')
		return T_DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c - 'a')
		return T_LETTER
	}
	return int(c)
}

func (l *CoffeeLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
