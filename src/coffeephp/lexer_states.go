package coffeephp

import (
	_ "fmt"
	"unicode"
)

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
		l.line++
		l.next()
		l.emit(T_NEWLINE)
		l.lastSpace = l.space
		l.space = 0
		return lexStartLine
	} else if c == '=' && l.peekMore(2) != '=' {
		l.next()
		l.emit(T_ASSIGN)
		return lexStart
	} else if c == '[' {
		l.next()
		l.emit(T_BRACKET_OPEN)
		return lexStart
	} else if c == ']' {
		l.next()
		l.emit(T_BRACKET_CLOSE)
		return lexStart
	} else if c == '"' || c == '\'' {
		return lexString
	} else if l.consumeIfMatch("//") {
		return lexOnelineComment
	} else if l.consumeIfMatch("/*") {
		return lexComment
	} else if l.emitIfMatch("if ", T_IF) {
		return lexStart
	} else if l.emitIfMatch("if else ", T_ELSEIF) {
		return lexStart
	} else if l.emitIfMatch("class ", T_CLASS) {
		return lexStart
	} else if l.emitIfMatch("for ", T_FOR) {
		return lexStart
	} else if l.emitIfMatch("foreach ", T_FOREACH) {
		return lexStart
	} else if l.emitIfMatch("echo ", T_ECHO) {
		return lexStart
	} else if l.emitIfMatch("does ", T_DOES) {
		return lexStart
	} else if l.emitIfMatch("is ", T_IS) {
		return lexStart
	} else if unicode.IsLetter(c) {
		return lexIdentifier
	} else if c == eof {
		l.emit(T_EOF)
		return nil
	} else {
		panic("unknown token.")
	}
	return nil
}

// Lex double quote string
func lexString(l *CoffeeLex) stateFn {
	var q rune = l.next() // the quote char

	l.ignore()

	var c rune
	for {
		c = l.next()
		if c == eof {
			panic("unexpected end of string.")
			break
		}
		// un-escaped string quote
		if c == q {
			if l.last() != '\\' {
				break
			}
		}
	}
	l.backup()
	l.emit(T_STRING)
	l.next()
	l.ignore() // ignore the quote rune
	return lexStart

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
	// consume the "/" char
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
