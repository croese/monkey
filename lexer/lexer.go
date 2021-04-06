package lexer

import (
	"github.com/croese/monkey/token"
)

type Lexer struct {
	input        string
	position     int // points to current char in input
	readPosition int // current reading position (after current char)
	ch           byte
	lineNumber   int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = l.newSingleCharToken(token.ASSIGN)
	case ';':
		tok = l.newSingleCharToken(token.SEMICOLON)
	case '(':
		tok = l.newSingleCharToken(token.LPAREN)
	case ')':
		tok = l.newSingleCharToken(token.RPAREN)
	case '{':
		tok = l.newSingleCharToken(token.LBRACE)
	case '}':
		tok = l.newSingleCharToken(token.RBRACE)
	case ',':
		tok = l.newSingleCharToken(token.COMMA)
	case '+':
		tok = l.newSingleCharToken(token.PLUS)
	case 0:
		tok = l.eofToken()
	}

	l.readChar()
	return tok
}

func (l *Lexer) newSingleCharToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(l.ch),
		Line:    l.lineNumber + 1,
		Column:  l.position + 1,
	}
}

func (l *Lexer) eofToken() token.Token {
	return token.Token{
		Type:    token.EOF,
		Literal: "",
		Line:    l.lineNumber + 1,
		Column:  l.position + 1,
	}
}
