package main

import (
	"fmt"
	"unicode"
)

type CoffeeLex struct {
	// the line
	s   string
	pos int
}

// returns a token
func (l *CoffeeLex) Lex(lval *CoffeeSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c - '0')
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c - 'a')
		return LETTER
	}
	return int(c)
}

func (l *CoffeeLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
