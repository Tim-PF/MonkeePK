package parser

import (
"github.com/Tim-PF/MonkeePK/lexer"
"github.com/Tim-PF/MonkeePK/ast"
"github.com/Tim-PF/MonkeePK/token"
)
// Exist of a lexer pointer current and next token
type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}
// Creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Get the next token and current token to work with
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() // in lexer not in parser. Not the same
	}
func (p *Parser) ParseProgram() *ast.Program {
	return nil
	}
