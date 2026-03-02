package parser

import (
	"foo/ast"
	"foo/tokens"
)

type bindingPower int

const (
	default_bp bindingPower = iota
	primary
	assignment
	unary
)

type (
	stmtHandler func(p *parser) ast.Stmt
	nudHandler  func(p *parser) ast.Expr
	ledHandler  func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

	stmtLookup map[tokens.TokenType]stmtHandler
	ledLookup  map[tokens.TokenType]ledHandler
	nudLookup  map[tokens.TokenType]nudHandler
	bpLookup   map[tokens.TokenType]bindingPower
)

var (
	stmtLu = stmtLookup{}
	bpLu   = bpLookup{}
	ledLu  = ledLookup{}
	nudLu  = nudLookup{}
)

func nud(kind tokens.TokenType, nudFn nudHandler) {
	bpLu[kind] = primary
	nudLu[kind] = nudFn
}

func led(kind tokens.TokenType, bp bindingPower, ledFn ledHandler) {
	bpLu[kind] = bp
	ledLu[kind] = ledFn
}

func stmt(kind tokens.TokenType, stmtFn stmtHandler) {
	bpLu[kind] = default_bp
	stmtLu[kind] = stmtFn
}

func createTokenLookups() {
	led(tokens.ASSIGNMENT, assignment, parseAssignmentExpr)

	nud(tokens.IDENT, parsePrimaryExpr)
	nud(tokens.STRING, parsePrimaryExpr)
	nud(tokens.BOOL, parsePrimaryExpr)
	nud(tokens.L_BRACE, parseArrayExpr)
	nud(tokens.L_CURL, parseObjectExpr)

	stmt(tokens.L_CURL, parseBlockStmt)
	stmt(tokens.SET, parseSetStmt)
}
