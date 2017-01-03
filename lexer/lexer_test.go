package lexer

import (
	"github.com/jcc333/monkey/token"
	"testing"
  "fmt"
)

func TestLexer(t *testing.T) {
	input := `let five = 5;
  let ten = 10f;
  let add = fn(x, y) {
    x + y;
  };
  let result = add(five, ten);

  !-/*5;
  5 < 10 > 5;

  if (5 < 10) {
    return true;
  } else {
    return false;
  }
  10 == 10;
  10 != 9;
  `
	tests := []token.Token {
		token.LET,
		{token.IDENT, "five"},
		{token.OPERATOR, "="},
		{token.INT, "5"},
		token.SEMICOLON,
		token.LET,
		{token.IDENT, "ten"},
		{token.OPERATOR, "="},
		{token.FLOAT, "10f"},
		token.SEMICOLON,
		token.LET,
		{token.IDENT, "add"},
		{token.OPERATOR, "="},
		token.FUNCTION,
		token.LPAREN,
		{token.IDENT, "x"},
		{token.OPERATOR, ","},
		{token.IDENT, "y"},
		token.RPAREN,
		token.LBRACE,
		{token.IDENT, "x"},
		{token.OPERATOR, "+"},
		{token.IDENT, "y"},
		token.SEMICOLON,
		token.RBRACE,
		token.SEMICOLON,
		token.LET,
		{token.IDENT, "result"},
		{token.OPERATOR, "="},
		{token.IDENT, "add"},
		token.LPAREN,
		{token.IDENT, "five"},
		{token.OPERATOR, ","},
		{token.IDENT, "ten"},
		token.RPAREN,
		token.SEMICOLON,
    {token.OPERATOR, "!-/*"},
		{token.INT, "5"},
		token.SEMICOLON,
		{token.INT, "5"},
    {token.OPERATOR, "<"},
		{token.INT, "10"},
    {token.OPERATOR, ">"},
		{token.INT, "5"},
		token.SEMICOLON,
		token.IF,
		token.LPAREN,
		{token.INT, "5"},
    {token.OPERATOR, "<"},
		{token.INT, "10"},
		token.RPAREN,
		token.LBRACE,
    token.RETURN,
    token.TRUE,
		token.SEMICOLON,
		token.RBRACE,
    token.ELSE,
		token.LBRACE,
    token.RETURN,
    token.FALSE,
		token.SEMICOLON,
		token.RBRACE,
		{token.INT, "10"},
    {token.OPERATOR, "=="},
		{token.INT, "10"},
		token.SEMICOLON,
		{token.INT, "10"},
    {token.OPERATOR, "!="},
		{token.INT, "9"},
		token.SEMICOLON,
		token.EOF,
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
    fmt.Printf("%v\n", tok)
		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.Type, tok.Type)
		}
		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.Literal, tok.Literal)
		}
	}
}
