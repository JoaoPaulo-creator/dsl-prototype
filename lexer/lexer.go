package lexer

import (
	"fmt"
	"foo/tokens"
	"unicode"
)

type lexer struct {
	input  string
	pos    int
	line   int
	Tokens []tokens.Token
}

func Tokenize(source string) []tokens.Token {
	lex := &lexer{
		input:  source,
		pos:    0,
		line:   1,
		Tokens: make([]tokens.Token, 0),
	}

	for !lex.atEof() {
		lex.scanToken()
	}

	lex.push(newToken(tokens.EOF, "EOF"))
	return lex.Tokens
}

func (l *lexer) scanToken() {
	// l.skipWhitespace()
	ch := l.peek()

	if unicode.IsSpace(rune(ch)) {
		l.skipWhitespace()
		return
	}

	// skipping comments
	if ch == '-' && l.peekNext() == '-' {
		l.skipComment()
		return
	}

	if unicode.IsLetter(rune(ch)) {
		l.advance()
		ident := l.scanIdentifier()
		l.push(newToken(tokens.LookupIdent(ident), ident))
		return
	}

	if ch == '"' {
		l.scanString()
		return
	}

	switch ch {
	case '(':
		l.advance()
		l.push(newToken(tokens.L_PAREN, "("))
		return
	case ')':
		l.advance()
		l.push(newToken(tokens.R_PAREN, ")"))
		return
	case '{':
		l.advance()
		l.push(newToken(tokens.L_CURL, "{"))
		return
	case '}':
		l.advance()
		l.push(newToken(tokens.R_CURL, "}"))
		return
	case '=':
		l.advance()
		l.push(newToken(tokens.ASSIGNMENT, "="))
		return
	case ',':
		l.advance()
		l.push(newToken(tokens.COMMA, ","))
		return
	case '[':
		l.advance()
		l.push(newToken(tokens.L_BRACE, "["))
		return
	case ']':
		l.advance()
		l.push(newToken(tokens.R_BRACE, "]"))
		return
	}

	panic(fmt.Sprintf("lexer error: unexpected character '%c' at position '%d'\n", ch, l.pos))
}

func (l *lexer) scanString() {

	start := l.pos
	l.advance()
	for !l.atEof() && l.peek() != '"' {
		l.advance()
	}

	l.advance()
	value := l.input[start:l.pos]
	l.push(newToken(tokens.STRING, value))
}

func (l *lexer) scanIdentifier() string {
	start := l.pos

	for !l.atEof() {
		ch := l.peek()
		if unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch)) || ch == '_' {
			l.advance()
		} else {
			break
		}
	}

	value := l.input[start:l.pos]
	l.push(newToken(tokens.IDENT, value))
	return l.input[start:l.pos]
}

func (l *lexer) scanBoolean() {
	start := l.pos
	l.advance()
	for !l.atEof() {
		l.advance()
	}

	l.advance()
	value := l.input[start:l.pos]
	l.push(newToken(tokens.BOOL, value))
}

func (l *lexer) skipWhitespace() {
	for l.peek() == ' ' || l.peek() == '\t' || l.peek() == '\n' || l.peek() == '\r' {
		l.advance()
	}
}

func (l *lexer) skipComment() {
	for !l.atEof() && l.peek() != '\n' {
		l.advance()
	}

	if !l.atEof() {
		l.advance()
		l.line++
	}
}

func (l *lexer) peek() byte {
	if l.atEof() {
		return 0
	}

	return l.input[l.pos]
}

func (l *lexer) peekNext() byte {
	if l.pos+1 >= len(l.input) {
		return 0
	}

	return l.input[l.pos+1]
}

func (l *lexer) advance() {
	l.pos++
}

func (l *lexer) push(token tokens.Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *lexer) atEof() bool {
	return l.pos >= len(l.input)
}

func newToken(kind tokens.TokenType, value string) tokens.Token {
	return tokens.Token{
		Type:    kind,
		Literal: value,
	}
}
