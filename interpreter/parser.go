package interpreter

import (
	"strconv"
)

type Priority = int

const (
	_ Priority = iota
	PriorityZero

	PriorityOne
	PriorityTwo
	PriorityThree

	PriorityInfinite
)

var tokenPriorities = map[TokenType]Priority{
	PlusToken:  PriorityOne,
	MinusToken: PriorityOne,

	SlashToken:    PriorityTwo,
	AsteriskToken: PriorityTwo,
}

type (
	prefixExprParser func() Expression
	infixExprParser  func(left Expression) Expression
	Parser           struct {
		lexer     *Lexer
		currToken Token
		nextToken Token

		prefixExprFns map[TokenType]prefixExprParser
		infixExprFns  map[TokenType]infixExprParser
	}
)

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer:         lexer,
		prefixExprFns: map[TokenType]prefixExprParser{},
		infixExprFns:  map[TokenType]infixExprParser{},
	}

	p.registerPrefixParser(IntegerToken, p.parseIntegerExpression)
	p.registerPrefixParser(MinusToken, p.parsePrefixOperatorExpression)

	infixTokens := []TokenType{PlusToken, MinusToken, SlashToken, AsteriskToken}
	for _, infixTokenType := range infixTokens {
		p.registerInfixParser(infixTokenType, p.parseInfixExpression)
	}

	p.advanceToken()
	p.advanceToken()

	return p
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{
		Expressions: []Expression{},
	}

	for !p.currTokenIs(EOFToken) {
		if expr := p.parseExpression(PriorityZero); expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}

		p.advanceToken()
	}

	return program
}

func (p *Parser) parseExpression(priority Priority) Expression {
	prefixParser, hasPrefixParser := p.prefixExprFns[p.currToken.Type]
	if !hasPrefixParser {
		return nil
	}

	leftExpression := prefixParser()

	for p.nextTokenPriority() > priority {
		infixParser, hasNextInfixParser := p.infixExprFns[p.nextToken.Type]
		if !hasNextInfixParser {
			return leftExpression
		}
		p.advanceToken()

		leftExpression = infixParser(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parsePrefixOperatorExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
	}

	p.advanceToken()

	expression.Right = p.parseExpression(PriorityThree)

	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.currToken,
		Left:     left,
		Operator: p.currToken.Value,
	}

	priority := p.currTokenPriority()
	p.advanceToken()
	expression.Right = p.parseExpression(priority)

	return expression
}

func (p *Parser) parseIntegerExpression() Expression {
	expression := &IntegerExpression{
		Token: p.currToken,
	}

	value, err := strconv.ParseInt(p.currToken.Value, 0, 64)
	if err != nil {
		// todo
		return nil
	}
	expression.Value = value

	return expression
}

func (p *Parser) registerPrefixParser(tokenType TokenType, fn prefixExprParser) {
	p.prefixExprFns[tokenType] = fn
}

func (p *Parser) registerInfixParser(tokenType TokenType, fn infixExprParser) {
	p.infixExprFns[tokenType] = fn
}

func (p *Parser) nextTokenPriority() Priority {
	if p, ok := tokenPriorities[p.nextToken.Type]; ok {
		return p
	}
	return PriorityZero
}

func (p *Parser) currTokenPriority() Priority {
	if p, ok := tokenPriorities[p.currToken.Type]; ok {
		return p
	}
	return PriorityZero
}

func (p *Parser) advanceToken() {
	p.currToken = p.nextToken
	p.nextToken = p.lexer.NextToken()
}

func (p *Parser) currTokenIs(tokenType TokenType) bool {
	return p.currToken.Type == tokenType
}
