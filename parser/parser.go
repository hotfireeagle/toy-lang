package parser

import (
	"jpg/ast"
	"jpg/lexer"
	"jpg/tokentype"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  *tokentype.Token
	peekToken *tokentype.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Body = []ast.Statement{}

	for !p.curToken.IsEof() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Body = append(program.Body, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokentype.LET:
	case tokentype.VAR:
	case tokentype.CONST:
		return p.parseVariableDeclaration()
	}
	return nil
}

// define = let、var、const
// define a; expected the const, because if we use const, then it must have the right side init value
// define a, b; we also can just use one statement to define multiple variable
// define a = 3; we can specify the right value
// TODO: copy acorn
func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	// 先创建一个Node节点，这个Node节点是statement类型的
	vd := &ast.VariableDeclaration{
		Kind:         p.curToken.Type,
		Declarations: make([]ast.Node, 0),
	}

	p.nextToken()
	// continue consume untile we meet whiteSpace or semi
	// TODO: 要不要eof?
	for !p.curToken.IsSemi() && !p.curToken.IsEnter() {
		declarator := &ast.VariableDeclarator{}

		// we use it setting the id
		p.parseVariableDeclarator(declarator, vd.Kind)

		p.nextToken()

		// if curToken is =, then consume it, so after the consume
		// the curToken will be the real value
		if p.eat(tokentype.EQ) {
			declarator.Init = p.parseMaybeAssign()
		}

		vd.Declarations = append(vd.Declarations, declarator)

		p.nextToken()
	}

	return vd
}

func (p *Parser) parseVariableDeclarator(decl *ast.VariableDeclarator, tokentype tokentype.TokenType) {
	decl.Id = p.parseBindingAtom()
}

// three pattern
// const [a, b] = [b, a]
// const {a, b} = {a: 1, b: 2}
// const c = 3
func (p *Parser) parseBindingAtom() *ast.Identifier {
	tok := p.curToken

	if tok.Type == tokentype.BRACKETL {
		// TODO:
		return nil
	} else if tok.Type == tokentype.BRACEL {
		// TODO:
		return nil
	} else if tok.Type == tokentype.IDENTIFIER {
		return &ast.Identifier{
			Token: tok,
			Name:  tok.Literal,
		}
	} else {
		panic("wrong token, expect to be identifier, or {, [")
	}
}

func (p *Parser) parseMaybeAssign() ast.Node {
	// TODO: 补全省略的那些代码
	left := p.parseMaybeConditional()
	return left
}

func (p *Parser) parseMaybeConditional() ast.Node {
	expr := p.parseExprOps()

	if p.eat(tokentype.QUESTION) {
		// TODO:
		return nil
	}

	return expr
}

func (p *Parser) parseExprOps() ast.Node {
	// TODO: 省略了很多代码
	return p.parseMaybeUnary()
}

func (p *Parser) parseMaybeUnary() ast.Node {
	// TODO: 省略了很多代码
	expr := p.parseExprSubscripts()
	return expr
}

func (p *Parser) parseExprSubscripts() ast.Node {
	expr := p.parseExprAtom()
	// TODO: 还有很多其它判断
	return expr
}

func (p *Parser) parseExprAtom() ast.Node {
	tokType := p.curToken.Type

	if tokType == tokentype.NUM {
		return p.parseLiteral()
	}

	return nil
}

func (p *Parser) parseLiteral() *ast.Literal {
	node := &ast.Literal{}
	node.Token = *p.curToken
	// TODO: 确定一下这里要不要调用那个p.nextToken()
	return node
}

// check the left-value is valid
// there patterns
// const [a, b] = [1, 2]
// const {c, d} = {c: 3, d: 4}
// const f = 3
// different pattern has different check method
func (p *Parser) checkLValPattern() {

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
		return false
	}
}

func (p *Parser) eat(t tokentype.TokenType) bool {
	if p.curTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
