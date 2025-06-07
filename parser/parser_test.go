package parser

import (
	"testing"

	"github.com/aaylward/goterp/ast"
	"github.com/aaylward/goterp/lexer"
)

func TestLetStatments(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foo = 949494;
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got %q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got %T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got %s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
