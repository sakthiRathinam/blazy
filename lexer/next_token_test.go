package Lexer

import (
	"testing"

	token "github.com/sakthiRathinam/blazy/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType || tt.expectedLiteral != tok.Literal {
			t.Fail()
		}
	}
}
