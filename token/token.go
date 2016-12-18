package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func Empty() Token {
  return EOF
}

var keywords = map[string]Token{
  "fn": FUNCTION,
  "let": LET,
  "if": IF,
  "else": ELSE,
  "true": TRUE,
  "false": FALSE,
  "return": RETURN,
}

func CheckKeyword(ident string) Token {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return Token{IDENT, ident}
}

func IntOrFloat(number string) Token {
  if strings.HasSuffix(number, "f") || strings.Contains(number, ".") {
    return Token{FLOAT, number}
  }
  return Token{INT, number}
}

const (
  EOF_T = "EOF"
	ILLEGAL = "ILLEGAL"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
  FLOAT = "FLOAT" // 1343.456

	// Operators
  OPERATOR  = "OPERATOR"
	EQ = "="
	PLUS   = "+"

  // Delimiters
  COMMA_T     = ","
  SEMICOLON_T = ";"
  LPAREN_T    = "("
  RPAREN_T    = ")"
  LBRACE_T    = "{"
  RBRACE_T    = "}"
  LBRACK_T    = "["
  RBRACK_T    = "]"

  // Keywords
  FUNCTION_T = "FUNCTION"
  LET_T      = "LET"
  IF_T       = "IF"
  ELSE_T     = "ELSE"
  TRUE_T     = "TRUE"
  FALSE_T    = "FALSE"
  RETURN_T   = "RETURN"
)

// Delimiters
var COMMA     = Token{",", ","}
var	SEMICOLON = Token{";", ";"}
var	LPAREN    = Token{"(", "("}
var	RPAREN    = Token{")", ")"}
var	LBRACE    = Token{"{", "{"}
var	RBRACE    = Token{"}", "}"}
var LBRACK    = Token{"[", "["}
var RBRACK    = Token{"]", "]"}
var	EOF       = Token{"EOF", ""}

// Keywords
var FUNCTION = Token{"FUNCTION", "fn"}
var LET      = Token{"LET", "let"}
var IF       = Token{"IF", "if"}
var ELSE     = Token{"ELSE", "else"}
var TRUE     = Token{"TRUE", "true"}
var FALSE    = Token{"FALSE", "false"}
var RETURN   = Token{"RETURN", "return"}
