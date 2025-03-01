package Lexer

import (
	"testing"

	token "github.com/sakthiRathinam/blazy/token"
)

type testToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {

	baseTestInput := "=+(){},;"
	baseTestTokens := []testToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	tests := []struct {
		testTokens []testToken
		input      string
	}{{testTokens: baseTestTokens, input: baseTestInput}}

	for _, test := range tests {
		l := NewLexer(test.input)
		for _, tt := range test.testTokens {
			tok := l.NextToken()
			if tok.Type != tt.expectedType || tt.expectedLiteral != tok.Literal {
				t.Fail()
			}
		}
	}
}
