package lexer

import "github.com/aaylward/goterp/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (next position to read)
	ch           byte // current character
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		if l.nextCharIs('=') {
			tok = newToken(token.EQUALS, "==")
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '!':
		if l.nextCharIs('=') {
			tok = newToken(token.NOT_EQUALS, "!=")
			l.readChar()
		} else {
			tok = newToken(token.NOT, string(l.ch))
		}
	case '>':
		if l.nextCharIs('=') {
			tok = newToken(token.GREATER_EQUALS, ">=")
			l.readChar()
		} else {
			tok = newToken(token.GREATER, string(l.ch))
		}
	case '<':
		if l.nextCharIs('=') {
			tok = newToken(token.LESSER_EQUALS, "<=")
			l.readChar()
		} else {
			tok = newToken(token.LESSER, string(l.ch))
		}
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case '[':
		tok = newToken(token.LBRACKET, string(l.ch))
	case ']':
		tok = newToken(token.RBRACKET, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '*':
		tok = newToken(token.MULTIPLY, string(l.ch))
	case '/':
		tok = newToken(token.DIVIDE, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func (l *Lexer) nextCharIs(ch byte) bool {
	return l.readPosition < len(l.input) && l.input[l.readPosition] == ch
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
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
