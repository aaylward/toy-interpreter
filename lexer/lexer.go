package lexer

import "github.com/aaylward/goterp/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (next position to read)
	ch           byte // current character
	line         int  // current line number
	column       int  // current column number
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

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.nextCharIs('=') {
			tok = l.makeTwoCharacterToken(token.EQUALS, "==")
		} else {
			tok = newToken(token.ASSIGN, string(l.ch), l.line, l.column)
		}
	case '!':
		if l.nextCharIs('=') {
			tok = l.makeTwoCharacterToken(token.NOT_EQUALS, "!=")
		} else {
			tok = newToken(token.NOT, string(l.ch), l.line, l.column)
		}
	case '>':
		if l.nextCharIs('=') {
			tok = l.makeTwoCharacterToken(token.GREATER_THAN_EQUALS, ">=")
		} else {
			tok = newToken(token.GREATER_THAN, string(l.ch), l.line, l.column)
		}
	case '<':
		if l.nextCharIs('=') {
			tok = l.makeTwoCharacterToken(token.LESS_THAN_EQUALS, "<=")
		} else {
			tok = newToken(token.LESS_THAN, string(l.ch), l.line, l.column)
		}
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch), l.line, l.column)
	case ',':
		tok = newToken(token.COMMA, string(l.ch), l.line, l.column)
	case '(':
		tok = newToken(token.LPAREN, string(l.ch), l.line, l.column)
	case ')':
		tok = newToken(token.RPAREN, string(l.ch), l.line, l.column)
	case '{':
		tok = newToken(token.LBRACE, string(l.ch), l.line, l.column)
	case '}':
		tok = newToken(token.RBRACE, string(l.ch), l.line, l.column)
	case '[':
		tok = newToken(token.LBRACKET, string(l.ch), l.line, l.column)
	case ']':
		tok = newToken(token.RBRACKET, string(l.ch), l.line, l.column)
	case '+':
		tok = newToken(token.PLUS, string(l.ch), l.line, l.column)
	case '-':
		tok = newToken(token.MINUS, string(l.ch), l.line, l.column)
	case '*':
		tok = newToken(token.MULTIPLY, string(l.ch), l.line, l.column)
	case '/':
		tok = newToken(token.DIVIDE, string(l.ch), l.line, l.column)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = -1
		tok.Column = -1
	default:
		if isLetter(l.ch) {
			tok.Line = l.line
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Line = l.line
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch), l.line, l.column)
		}
	}

	l.readChar()
	l.column += 1
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) nextCharIs(ch byte) bool {
	return l.readPosition < len(l.input) && l.input[l.readPosition] == ch
}

func (l *Lexer) makeTwoCharacterToken(tokenType token.TokenType, literal string) token.Token {
	tok := newToken(tokenType, literal, l.line, l.column)
	l.readChar()
	l.column += 1
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber curently only supports non-negative ints
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, literal string, line int, column int) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Line: line, Column: column}
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
