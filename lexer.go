package main

import (
	"fmt"
	"unicode"
)

type CoffeeLexToken struct {
	// the line
	s   string
	pos int
}

type TokenType int

const (
	TokenDot = iota
	TokenEOF
	TokenIf
	TokenElse
	TokenElseIf
	TokenDigit
	TokenIdentifier
	TokenForeach
)

type LexItem struct {
	Val string
	Typ TokenType
}

func (self *LexItem) String() string {
	switch self.Typ {
	case TokenEOF:
		return "EOF"
	}
	return fmt.Sprintf("%q", self.Val)
}

// returns a token
func (l *CoffeeLexToken) Lex(lval *CoffeeSymType) int {
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

func (l *CoffeeLexToken) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
