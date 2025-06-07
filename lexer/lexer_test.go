package lexer

import (
	"testing"

	"github.com/aaylward/goterp/token"
)

func TestNextTokenWithBasicOperatorsAndDelimiters(t *testing.T) {
	input := "=+(){}[],;!=/==!*->>=<<="

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.ASSIGN, "=", 0, 0},
		{token.PLUS, "+", 0, 1},
		{token.LPAREN, "(", 0, 2},
		{token.RPAREN, ")", 0, 3},
		{token.LBRACE, "{", 0, 4},
		{token.RBRACE, "}", 0, 5},
		{token.LBRACKET, "[", 0, 6},
		{token.RBRACKET, "]", 0, 7},
		{token.COMMA, ",", 0, 8},
		{token.SEMICOLON, ";", 0, 9},
		{token.NOT_EQUALS, "!=", 0, 10},
		{token.DIVIDE, "/", 0, 12},
		{token.EQUALS, "==", 0, 13},
		{token.NOT, "!", 0, 15},
		{token.MULTIPLY, "*", 0, 16},
		{token.MINUS, "-", 0, 17},
		{token.GREATER_THAN, ">", 0, 18},
		{token.GREATER_THAN_EQUALS, ">=", 0, 19},
		{token.LESS_THAN, "<", 0, 21},
		{token.LESS_THAN_EQUALS, "<=", 0, 22},
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

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - wrong line number. expected %d got %d", i, tt.expectedLine, tok.Line)
		}

		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - wrong column number. expected %d got %d", i, tt.expectedColumn, tok.Column)
		}
	}
}

func TestNextTokenWithSimpleCodeExample(t *testing.T) {
	input := `
	let five = 5;
	let ten  = 10;

	let add = fn(x, y) {
		return x + y;
	};

	let result = add(five, ten);

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
// 	if (5 < 10) {
// return true;
// } else {
// return false;
// }

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
