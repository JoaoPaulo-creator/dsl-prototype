package parser

import (
	"foo/ast"
	"foo/tokens"
)

type bindingPower int

const (
	default_bp bindingPower = iota
)

type (
	stmtHandler func(p *parser) ast.Stmt
	stmtLookup  map[tokens.TokenType]stmtHandler
	bpLookUp    map[tokens.TokenType]bindingPower
)

var (
	stmtLu = stmtLookup{}
	bpLu   = bpLookUp{}
)

func stmt(kind tokens.TokenType, stmtFn stmtHandler) {
	bpLu[kind] = default_bp
	stmtLu[kind] = stmtFn
}

func createTokenLookups() {
	stmt(tokens.SET, parseSetStmt)
}
