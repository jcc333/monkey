package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func Empty() Token {
  return Token{Type: EOF, Literal:"" }
}

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "if": IF,
  "else": ELSE,
  "true": TRUE,
  "false": FALSE,
  "return": RETURN,
}

func CheckKeyword(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}

func IntOrFloat(number string) TokenType {
  if strings.HasSuffix(number, "f") || strings.Contains(number, ".") {
    return FLOAT
  }
  return INT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
  FLOAT = "FLOAT" // 1343.456

	// Operators
  OPERATOR  = "OPERATOR"
	EQ = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
  LBRACK    = "["
  RBRACK    = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
  IF       = "IF"
  ELSE     = "ELSE"
  TRUE     = "TRUE"
  FALSE    = "FALSE"
  RETURN   = "RETURN"
)
