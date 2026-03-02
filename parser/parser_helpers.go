package parser

import (
	"foo/tokens"
)

func (p *parser) currentToken() tokens.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() tokens.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != tokens.EOF
}

func (p *parser) currentTokenKind() tokens.TokenType {
	return p.tokens[p.pos].Type
}

func (p *parser) nextToken() tokens.Token {
	return p.tokens[p.pos+1]
}

func (p *parser) previousToken() tokens.Token {
	return p.tokens[p.pos-1]
}

func (p *parser) expectError(expectedKind tokens.TokenType, err any) tokens.Token {
	token := p.currentToken()
	kind := token.Type

	if kind != expectedKind {
		if err == nil {
			panic(err)
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expectedKind tokens.TokenType) tokens.Token {
	return p.expectError(expectedKind, nil)
}
