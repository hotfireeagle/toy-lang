package parser

import (
	"fmt"
	"strconv"
	"toy/ast"
	"toy/lexer"
	"toy/tokentype"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGRATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	curToken       *tokentype.Token
	peekToken      *tokentype.Token
	errors         []string
	prefixParseFns map[tokentype.TokenType]prefixParseFn
	infixParseFns  map[tokentype.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		errors:         []string{},
		prefixParseFns: make(map[tokentype.TokenType]prefixParseFn),
		infixParseFns:  make(map[tokentype.TokenType]infixParseFn),
	}

	p.registerPrefix(tokentype.IDENTIFIER, p.parseIdentfier)
	p.registerPrefix(tokentype.NUM, p.parseIntegerLiteral)
	p.registerPrefix(tokentype.NOT, p.parsePrefixExpression)
	p.registerPrefix(tokentype.MIN, p.parsePrefixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokenType tokentype.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType tokentype.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseIdentfier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != tokentype.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokentype.LET:
		return p.parseLetStatement()
	case tokentype.RETURN:
		return p.parseReturnStatement()
	default:
		if p.curToken.IsEnter() {
			return nil
		}
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(tokentype.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(tokentype.EQ) {
		return nil
	}

	for !p.curToken.IsSemi() {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: *p.curToken,
	}
	p.nextToken()
	for !p.curTokenIs(tokentype.SEMI) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: *p.curToken,
	}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(tokentype.SEMI) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) noPrefixParseFnError(t *tokentype.Token) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t.Literal)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
		return nil
	}
	lextExp := prefix()
	return lextExp
}

func (p *Parser) curTokenIs(t tokentype.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t tokentype.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t tokentype.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t tokentype.TokenType) {
	msg := fmt.Sprintf("expected next token to be %v, got %v instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// TODO: 需要区别浮点数
func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{
		Token: *p.curToken,
	}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("count not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	il.Value = val
	return il
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    *p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
