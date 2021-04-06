package lexer

import (
	"testing"

	"github.com/croese/monkey/token"
	"github.com/stretchr/testify/require"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.ASSIGN, "=", 1, 1},
		{token.PLUS, "+", 1, 2},
		{token.LPAREN, "(", 1, 3},
		{token.RPAREN, ")", 1, 4},
		{token.LBRACE, "{", 1, 5},
		{token.RBRACE, "}", 1, 6},
		{token.COMMA, ",", 1, 7},
		{token.SEMICOLON, ";", 1, 8},
		{token.EOF, "", 1, 9},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		require.Equal(t, tt.expectedType, tok.Type, "in tests[%d]", i)
		require.Equal(t, tt.expectedLiteral, tok.Literal, "in tests[%d]", i)
		require.Equal(t, tt.expectedLine, tok.Line, "in tests[%d]", i)
		require.Equal(t, tt.expectedColumn, tok.Column, "in tests[%d]", i)
	}
}
