package lexer

import (
  "github.com/jcc333/monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	current  token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, readPosition: 0, ch: 0, current: token.Empty()}
  l.readChar()
  l.next()
	return l
}

func (l *Lexer) NextToken() token.Token {
  tok := l.Current()
  l.next()
  return tok
}

func (l *Lexer) Current() token.Token {
	return l.current
}

func (l *Lexer) next() bool {
	l.current = token.Empty()

  l.skipWhitespace()
  if !(l.readWord() || l.readNumber() || l.readOperator()) {
    switch l.ch {
    case ';':
      l.current = token.SEMICOLON
    case '(':
      l.current = token.LPAREN
    case ')':
      l.current = token.RPAREN
    case ',':
      l.current = token.COMMA
    case '{':
      l.current = token.LBRACE
    case '}':
      l.current = token.RBRACE
    case '[':
      l.current = token.LBRACK
    case ']':
      l.current = token.RBRACK
    case 0:
      return false
    default:
      l.current = token.Token{token.ILLEGAL, string(l.ch)}
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
    } else if l.ch == '.' {
      if isDigit(l.peekChar()) {
        l.readChar()
        for isDigit(l.ch) {
          l.readChar()
        }
        stop = l.position
      }
    }
    literal := l.input[start:stop]
    l.current = token.IntOrFloat(literal)
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
    l.current = token.CheckKeyword(literal)
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
    l.current = token.Token{token.OPERATOR, literal}
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
}

func (l *Lexer) peekChar() byte {
  if l.readPosition >= len(l.input) {
    return 0
  } else {
    return l.input[l.readPosition]
  }
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}
