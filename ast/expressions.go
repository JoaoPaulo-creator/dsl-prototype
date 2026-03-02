package ast

import "foo/tokens"

type SymbolExpr struct {
	Value string
}

func (node SymbolExpr) expr() {}

type AssignmentExpr struct {
	Assignee Expr
	Operator tokens.Token
	Value    Expr
}

func (node AssignmentExpr) expr() {}

type StringExpr struct {
	Value string
}

func (node StringExpr) expr() {}

type BoolExpr struct {
	Value bool
}

func (node BoolExpr) expr() {}

type ArrayExpr struct {
	Elements []Expr
}

func (node ArrayExpr) expr() {}

type ObjectExpr struct {
	Properties map[string]Expr
}

func (node ObjectExpr) expr() {}
