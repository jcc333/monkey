package ast

import "fmt"

type Node interface {
  literal() string
}

type Program []Node

type Identifier string

func (id Identifier) literal {
  return id
}

type Statement interface {
  Node
}

type Expression interface {
  Node
}

type Program struct {
  statements []Statement
}

type Let struct {
  Name  *Identifier
  Value Expression
}

func (let *Let)eval(env Env) Expression {
  env[let.Name] := let.Value.eval(env)

}

func (let *Let) literal() string {
  return fmt.Sprintf("let %s = %s;", let.Name, let.Value.literal())
}

type Int struct {
  Value int32
}

func (n *Int) eval() Expression {
  return n
}

func (n *Int) literal() string {
  return fmt.Sprintf("%d", n.Value)
}

type Float struct {
  Value float32
}

func (f *Float) literal() string {
  return fmt.Sprintf("%f", f.Value)
}

type String struct {
  Value string
}

func (f *Float) literal() string {
  return fmt.Sprintf("%f", f.Value)
}

type Function struct {
  Args []Identifier
  Body []Statement
}

func (f *Function) literal() string {
  lit := "fn("
  for arg in f.Args {
    lit += arg.literal()
  }
}


type Hash map[Identifier]Expression

type Array []Expression

type Return struct {
  Value Expression
}

func (r *Return) eval(env Env) {
  return r.Value
}

type Call struct {
  Func Identifier
  Args []Expression
}

type Bool struct {
  Value bool
}

func (b *Bool) literal() string {
  if b.Value {
    return "true"
  } else {
    return "false"
  }
}
