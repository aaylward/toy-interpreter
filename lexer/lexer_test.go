package lexer

import (
	"testing"
	"github.com/aaylward/goterp/token"
)

func TestNotEquals(t *testing.T) {
	input := "!="
	l := NewLexer(input)
	tok := l.NextToken()
	if tok.Type != token.NOT_EQUALS {
		t.Fatalf("expected type NOT_EQUALS, got %s", tok.Type)
	}
	if tok.Literal != "!=" {
		t.Fatalf("expected literal !=, got %s", tok.Literal)
	}
}

func TestNextToken(t *testing.T) {
	input := "=+(){}[],;!=/==!*->>=<<="

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.NOT_EQUALS, "!="},
		{token.DIVIDE, "/"},
		{token.EQUALS, "=="},
		{token.NOT, "!"},
		{token.MULTIPLY, "*"},
		{token.MINUS, "-"},
		{token.GREATER, ">"},
		{token.GREATER_EQUALS, ">="},
		{token.LESSER, "<"},
		{token.LESSER_EQUALS, "<="},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong TokenType. expected %s got %s", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong Literal. expected %s got %s", i, tt.expectedLiteral, tok.Literal)
		}
	}
}