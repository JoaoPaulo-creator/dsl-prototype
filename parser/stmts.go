package parser

import (
	"fmt"
	"foo/ast"
	"foo/tokens"
)

func parseStmt(p *parser) ast.Stmt {
	stmtFn, exists := stmtLu[p.currentTokenKind()]
	if exists {
		return stmtFn(p)
	}

	expression := parseExpr(p, default_bp)

	return ast.ExpressionStmt{
		ExpressionStmt: expression,
	}
}

func parseBlockStmt(p *parser) ast.Stmt {
	p.expect(tokens.L_CURL)
	body := []ast.Stmt{}
	for p.hasTokens() && p.currentTokenKind() != tokens.R_CURL {
		body = append(body, parseStmt(p))
	}

	p.expect(tokens.R_CURL)
	return ast.BlockStmt{
		Body: body,
	}
}

func parseSetStmt(p *parser) ast.Stmt {
	var properties = map[string]ast.StructProperty{}

	p.expect(tokens.SET)
	p.expect(tokens.L_CURL)

	for p.hasTokens() && p.currentTokenKind() != tokens.R_CURL {
		propertyName := p.advance().Literal
		_, exists := properties[propertyName]
		if exists {
			panic(fmt.Sprintf("property %s has already been defined inside struct declaration", propertyName))
		}

		p.expect(tokens.ASSIGNMENT)
		properyValue := parseExpr(p, default_bp)
		properties[propertyName] = ast.StructProperty{
			AssignedValue: properyValue,
		}

		p.expect(tokens.COMMA)
	}

	p.expect(tokens.R_CURL)
	return ast.StructDeclStmt{
		StructName: "set",
		Properties: properties,
	}
}
