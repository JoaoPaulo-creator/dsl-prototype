package parser

import (
	"foo/ast"
	"foo/tokens"
)

type parser struct {
	tokens []tokens.Token
	pos    int
}

func createParser(tokens []tokens.Token) *parser {
	createTokenLookups()

	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []tokens.Token) ast.BlockStmt {
	p := createParser(tokens)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() {
		body = append(body, parseStmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}
