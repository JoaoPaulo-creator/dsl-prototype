package parser

import (
	"foo/ast"
	"foo/tokens"
)

func parseStmt(p *parser) ast.Stmt {
	return ast.BlockStmt{}
}

func parseSetStmt(p *parser) ast.Stmt {
	var properties = map[string]ast.StructProperty{}
	structName := p.expect(tokens.IDENT).Literal
	p.expect(tokens.L_CURL)

	p.expect(tokens.R_CURL)

	return ast.StructDeclStmt{
		StructName: structName,
		Properties: properties,
	}
}
