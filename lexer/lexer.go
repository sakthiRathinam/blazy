package Lexer

import (
	token "github.com/sakthiRathinam/blazy/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	lexerPtr := &Lexer{input: input, position: 0, readPosition: 0}
	lexerPtr.readChar()
	return lexerPtr
}

func (l *Lexer) readChar() {
	if (l.readPosition) >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
func LookUPIdent(ident string) token.TokenType {

	tokenType, isKeyword := token.Keywords[ident]
	if isKeyword {
		return tokenType
	}
	return token.IDENT
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpaces()
	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			identifier := l.ParseIdentifier()
			tok.Literal = identifier
			tok.Type = LookUPIdent(identifier)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.ReadNum()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) ParseIdentifier() string {
	start_position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]
}
func (l *Lexer) ReadNum() string {
	start_position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]

}
func (l *Lexer) skipWhiteSpaces() {
	if l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}
