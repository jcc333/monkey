package lexer

import (
  "github.com/jcc333/monkey/token"
  "fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	accumulator  token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, readPosition: 0, ch: 0, accumulator: token.Empty()}
  fmt.Printf("Making lexer with input: `%s`\n", input)
  l.readChar()
  l.next()
  fmt.Printf("at time of lexer creation: %v\n", l.accumulator)
	return l
}

func (l *Lexer) NextToken() token.Token {
  tok := l.current()
  l.next()
  return tok
}

func (l *Lexer) current() token.Token {
  fmt.Printf("current: %v\n", l.accumulator)
	return l.accumulator
}

func (l *Lexer) next() bool {
	l.accumulator = token.Empty()

  l.skipWhitespace()
  if !(l.readWord() || l.readNumber() || l.readOperator()) {
    switch l.ch {
    case ';':
      l.accumulator = token.Token{token.SEMICOLON, string(l.ch)}
    case '(':
      l.accumulator = token.Token{token.LPAREN, string(l.ch)}
    case ')':
      l.accumulator = token.Token{token.RPAREN, string(l.ch)}
    case ',':
      l.accumulator = token.Token{token.COMMA, string(l.ch)}
    case '{':
      l.accumulator = token.Token{token.LBRACE, string(l.ch)}
    case '}':
      l.accumulator = token.Token{token.RBRACE, string(l.ch)}
    case '[':
      l.accumulator = token.Token{token.LBRACK, string(l.ch)}
    case ']':
      l.accumulator = token.Token{token.RBRACK, string(l.ch)}
    case 0:
      return false
    default:
      l.accumulator = token.Token{token.ILLEGAL, string(l.ch)}
      return false
    }
    l.readChar()
  }
	return true
}

func (l *Lexer) readNumber() bool {
  if isDigit(l.ch) {
    start := l.position
    for isDigit(l.ch) {
      l.readChar()
    }
    stop := l.position
    if l.ch == 'f' {
      stop += 1
      l.readChar()
    }
    literal := l.input[start:stop]
    l.accumulator = token.Token{token.IntOrFloat(literal), literal}
    fmt.Printf("Read Number: %v\n", l.accumulator)
    return true
  }
  return false
}

func (l *Lexer) readWord() bool {
  if isAlpha(l.ch) {
    start := l.position
    for isAlpha(l.ch) || isDigit(l.ch) {
      l.readChar()
    }
    stop := l.position
    literal := l.input[start:stop]
    l.accumulator = token.Token{token.CheckKeyword(literal), literal}
    fmt.Printf("Read Word: %v\n", l.accumulator)
    return true
  }
  return false
}

func isOpChar(c byte) bool {
  ops := []byte{'.', ',', '<', '>', '=', '=', '!', '@', '#', '$', '%', '^', '&', '*', '+', '-', '|', '/', '?'}
  for _, o := range ops {
    if o == c {
      return true
    }
  }
  return false
}

func (l *Lexer) readOperator() bool {
  if isOpChar(l.ch) {
    start := l.position
    for isOpChar(l.ch) {
      l.readChar()
    }
    stop := l.position
    literal := l.input[start:stop]
    l.accumulator = token.Token{token.OPERATOR, literal}
    fmt.Printf("Read Op: %v\n", l.accumulator)
    return true
  }
  return false
}

func isAlpha(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
  fmt.Printf("current char: %s\n", string(l.ch))
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}
