package parser

import (
	"fmt"
	"foo/ast"
	"foo/tokens"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	tokenKind := p.currentTokenKind()
	nudFn, exists := nudLu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("nud handler expected for token %s\n", tokens.TokenKindString(tokenKind)))
	}

	left := nudFn(p)
	for bpLu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		ledFn, exists := ledLu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("led handler expected for token %s\n", tokens.TokenKindString(tokenKind)))
		}

		left = ledFn(p, left, bpLu[p.currentTokenKind()])
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentTokenKind() {

	case tokens.IDENT:
		return ast.SymbolExpr{Value: p.advance().Literal}
	case tokens.STRING:
		return ast.StringExpr{Value: p.advance().Literal}
	case tokens.BOOL:
		return ast.BoolExpr{Value: p.advance().Literal == "true"}
	default:
		panic(fmt.Sprintf("cannot create primary expression from %s\n", tokens.TokenKindString(p.currentTokenKind())))
	}
}

func parseArrayExpr(p *parser) ast.Expr {
	p.expect(tokens.L_BRACE) // consume '['
	elements := []ast.Expr{}

	for p.hasTokens() && p.currentTokenKind() != tokens.R_BRACE {
		elements = append(elements, parseExpr(p, default_bp))
		if p.currentTokenKind() == tokens.COMMA {
			p.advance() // consume ','
		}
	}

	p.expect(tokens.R_BRACE) // consume ']'
	return ast.ArrayExpr{Elements: elements}
}

func parseAssignmentExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operator := p.advance()
	rhs := parseExpr(p, bp)
	return ast.AssignmentExpr{
		Operator: operator,
		Value:    rhs,
		Assignee: left,
	}
}

func parseObjectExpr(p *parser) ast.Expr {
	p.expect(tokens.L_CURL) // consume '{'
	properties := map[string]ast.Expr{}

	for p.hasTokens() && p.currentTokenKind() != tokens.R_CURL {
		key := p.advance().Literal
		p.expect(tokens.ASSIGNMENT)
		value := parseExpr(p, default_bp)
		properties[key] = value

		if p.currentTokenKind() == tokens.COMMA {
			p.advance() // consume ','
		}
	}

	p.expect(tokens.R_CURL) // consume '}'
	return ast.ObjectExpr{Properties: properties}
}
