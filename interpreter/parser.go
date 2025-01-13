package interpreter

type Parser struct {
	lexer     *Lexer
	currToken Token
	nextToken Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	p.advanceTokens()
	p.advanceTokens()

	return p
}

func (p *Parser) advanceTokens() {
	p.currToken = p.nextToken
	p.nextToken = p.lexer.NextToken()
}
