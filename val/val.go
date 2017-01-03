package val

import (
  "fmt"
  "reflect"
)

type MonkeyVal interface {
  display()     string
  eq(*MonkeyVal) bool
}

//
type Bool bool

func (b Bool) eq(m *MonkeyVal) bool {
  mb, ok := m.(Bool)
  if ok {
    mb == b
  }
  return false
}

func (b Bool) display() string {
  if b {
    return "true"
  } else {
    return "false"
  }
}

//
type Int int64

func (n Int) eq(m *MonkeyVal) bool {
  mn, ok := m.(Int)
  if ok {
    return mn == n
  }
}

func (n Int) display() string {
  format fmt.Sprintf("%d", n)
}

//
type Float float32

func (f Float) eq(m *MonkeyVal) bool {
  mf, ok := m.(Float)
  if ok {
    return f == mf
  }
  return false
}

func (f Float) display() string {
  format fmt.Sprintf("%f", f)
}

//
type Hash map[MonkeyVal]MonkeyVal

func (h Hash) eq(m *MonkeyVal) bool {
  mh, ok = m.(Hash)
  if ok {
    return reflect.DeepEqual(m, ma)
  }
  return false
}

func (h Hash) display() string {
  r := "{"
  for k, v := range h {
    r += k.display() + " : " + v.display()
  }
  r += "}"
  return r
}

//
type Array []MonkeyVal

func (a Array) eq(m *MonkeyVal) bool {
  ma, ok := m.(Array)
  if ok {
    return reflect.DeepEqual(m, ma)
  }
  return false
}

func (a Array) display() string {
  r := "["
  for v := range a {
    r += v.display()
  }
  r += "]"
  return r
}

//
type Function struct {
  Args []ast.Identifier
  Body []ast.Statement
  env  map[ast.Identifier]MonkeyVal
}

func (f Function) eq(m *MonkeyVal) {
  //structural, not extensional equality
  fm, ok == m.(Function)
  if ok {
    if fm.Args == f.Args && fm.Body == f.Body {
      for k, v := range f.env {
        fm.env[k] == v
      }
      for k, v := range fm.env {
      }
    }
  }
  return false
}

func (f Function) display() string {
  return fmt.Sprintf("<function of %d arguments>", len(f.Args))
}

//
type String string

func (s String) eq(m *MonkeyVal) {
  ms, ok := m.(String)
  if ok {
    return  && ms == s
  }
  return false
}

func (s String) display() string {
  return fmt.Sprintf("%q", s)
}
