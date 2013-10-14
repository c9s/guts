package gutscript

import (
	"fmt"
	"unicode"
)

var LexKeywords = map[string]int{
	"if":      T_IF,
	"class":   T_CLASS,
	"for":     T_FOR,
	"foreach": T_FOREACH,
	"else":    T_ELSE,
	"elseif":  T_ELSEIF,
	"echo":    T_ECHO,
	"does":    T_DOES,
	"is":      T_IS,
	"return":  T_RETURN,
}

func (l *GutsLex) emitIfKeywordMatches() bool {
	l.remember()
	for keyword, typ := range LexKeywords {
		var match bool = true
	next_keyword:
		for _, sc := range keyword {
			c := l.next()
			if c == eof {
				match = false
				break next_keyword
			}
			if sc != c {
				match = false
				break next_keyword
			}
		}
		if match {
			c := l.next()
			if c == '\n' || c == eof || c == ' ' || c == '\t' || unicode.IsSymbol(c) {
				l.backup()
				l.emit(TokenType(typ))
				return true
			}
		}
		l.rollback()
	}
	return false
}

func lexStartLine(l *GutsLex) stateFn {
	var c rune = l.peek()
	if c == ' ' || c == '\t' {
		return lexIndentSpaces
	}
	// default start state.
	return lexStart
}

func lexStart(l *GutsLex) stateFn {
	var c rune = l.peek()
	if unicode.IsDigit(c) {
		return lexNumber
	} else if l.emitIfMatch("->", T_FUNCTION_GLYPH) {
		return lexStart
	} else if l.emitIfMatch("::", T_FUNCTION_PROTOTYPE) {
		return lexStart
	} else if l.consumeIfMatch("//") {
		return lexOnelineComment
	} else if l.consumeIfMatch("/*") {
		return lexComment
	} else if c == '-' {
		l.next()
		l.emit(TokenType(c))
		return lexStart
	} else if c == ' ' || c == '\t' {
		// return lexSpaces
		return lexIgnoreSpaces
	} else if c == '\n' || c == '\r' {
		// if there is a new line, check the next line is indent or outdent,
		// if there is no spaces/indent in the next line, then it should be outdent.
		l.line++
		c = l.next()

		// skip multiple newline at one time
		for c == '\n' {
			c = l.next()
		}

		// c = l.peek()
		if c == eof {
			// if we're in an indent block, and it's the end of file.
			// we should treat the newline as a block end.
			if l.space > 0 {
				l.emit(T_OUTDENT)
			} else {
				l.ignore()
			}
			return nil
		}

		// reset space info
		l.lastSpace = l.space
		l.space = 0

		// if there are one or more space in the next line
		// consume the spaces and guess it's
		// indent/outdent/newline
		for c == ' ' || c == '\t' {
			l.space++
			c = l.next()
		}

		l.backup()
		if l.space == l.lastSpace {
			l.emit(T_NEWLINE)
		} else if l.space < l.lastSpace {
			l.emit(T_OUTDENT)
		} else if l.space > l.lastSpace {
			l.emit(T_INDENT)
		}
		return lexStartLine
	} else if c == '=' && l.peekMore(2) != '=' {
		l.next()
		l.emit(TokenType(c))
		return lexStart
	} else if l.accept("+-*|&[]{}()<>,") {
		l.emit(TokenType(c))
		return lexStart
	} else if l.lastTokenType == T_NUMBER && l.emitIfMatch("..", T_RANGE_OPERATOR) {
		return lexStart
	} else if c == '"' || c == '\'' {
		return lexString
	} else if l.emitIfKeywordMatches() {
		return lexStart
	} else if unicode.IsLetter(c) {
		return lexIdentifier
	} else if c == eof {
		// l.emit(T_EOF)
		return nil
	} else {
		panic(fmt.Errorf("unknown token %c\n", c))
		return nil
	}
	return nil
}

// Lex double quote string
func lexString(l *GutsLex) stateFn {
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

func lexComment(l *GutsLex) stateFn {
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

func lexOnelineComment(l *GutsLex) stateFn {
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

func lexIdentifier(l *GutsLex) stateFn {
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

func lexIgnoreSpaces(l *GutsLex) stateFn {
	var c rune
	for {
		c = l.next()
		if c == eof {
			l.ignore()
			return nil
		}
		if c != ' ' {
			break
		}
	}
	l.backup()
	l.ignore()
	return lexStart
}

func lexIndentSpaces(l *GutsLex) stateFn {
	l.space = 0
	var c rune
	for {
		c = l.next()
		if c == ' ' || c == '\t' {
			l.space++
		} else {
			break
		}
	}
	l.backup()
	l.ignore() // simply ignore string,
	return lexStart
}

func lexSpaces(l *GutsLex) stateFn {
	for c := l.next(); c == ' ' || c == '\t'; {
	}
	l.backup()
	l.emit(T_SPACE)
	return lexStart
}

func lexFloating(l *GutsLex) stateFn {
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

func lexNumber(l *GutsLex) stateFn {
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
