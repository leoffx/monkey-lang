package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	c            byte // char only supports UTF-8
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.c = 0
	} else {
		l.c = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch {
	case l.c == '=':
		tok = newToken(token.ASSIGN, l.c)
	case l.c == ';':
		tok = newToken(token.SEMICOLON, l.c)
	case l.c == '(':
		tok = newToken(token.LPAREN, l.c)
	case l.c == ')':
		tok = newToken(token.RPAREN, l.c)
	case l.c == ',':
		tok = newToken(token.COMMA, l.c)
	case l.c == '+':
		tok = newToken(token.PLUS, l.c)
	case l.c == '{':
		tok = newToken(token.LBRACE, l.c)
	case l.c == '}':
		tok = newToken(token.RBRACE, l.c)
	case l.c == 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case isLetter(l.c):
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	case isDigit(l.c):
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return tok
	default:
		tok = newToken(token.ILLEGAL, l.c)
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.c == ' ' || l.c == '\t' || l.c == '\n' || l.c == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, c byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(c)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.c) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.c) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}
