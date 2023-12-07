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
	switch l.c {
	case '=':
		tok = newToken(token.ASSIGN, l.c)
	case ';':
		tok = newToken(token.SEMICOLON, l.c)
	case '(':
		tok = newToken(token.LPAREN, l.c)
	case ')':
		tok = newToken(token.RPAREN, l.c)
	case ',':
		tok = newToken(token.COMMA, l.c)
	case '+':
		tok = newToken(token.PLUS, l.c)
	case '{':
		tok = newToken(token.LBRACE, l.c)
	case '}':
		tok = newToken(token.RBRACE, l.c)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, c byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(c)}
}
