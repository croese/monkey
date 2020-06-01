package lexer

import "github.com/croese/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
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
	l.readPosition++
	l.column++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	col := l.column

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch, l.line, col)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line, col)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line, col)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line, col)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line, col)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line, col)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line, col)
	case '!':
		tok = newToken(token.BANG, l.ch, l.line, col)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.line, col)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.line, col)
	case '<':
		tok = newToken(token.LT, l.ch, l.line, col)
	case '>':
		tok = newToken(token.GT, l.ch, l.line, col)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line, col)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line, col)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
		tok.Column = col
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.line
			tok.Column = col
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Line = l.line
			tok.Column = col
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, col)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' ||
		l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte, line, col int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Column:  col,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
