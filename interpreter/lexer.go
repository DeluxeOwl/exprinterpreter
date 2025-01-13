package interpreter

import (
	"slices"
)

type TokenType string

const (
	PlusToken             TokenType = "PlusToken"
	MinusToken            TokenType = "MinusToken"
	AsteriskToken         TokenType = "AsteriskToken"
	SlashToken            TokenType = "SlashToken"
	IntegerToken          TokenType = "IntegerToken"
	LeftParenthesisToken  TokenType = "LeftParenthesisToken"
	RightParenthesisToken TokenType = "RightParenthesisToken"
	EOFToken              TokenType = "EOFToken"
	IllegalToken          TokenType = "IllegalToken"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	sourceCode  string
	currCharPos int
	nextCharPos int
	currChar    byte
}

func NewLexer(sourceCode string) *Lexer {
	l := &Lexer{
		sourceCode: sourceCode,
	}
	l.nextChar()
	return l
}

func (l *Lexer) nextChar() {
	if l.nextCharPos >= len(l.sourceCode) {
		l.currChar = 0
	} else {
		l.currChar = l.sourceCode[l.nextCharPos]
	}
	l.currCharPos = l.nextCharPos
	l.nextCharPos++
}

func (l *Lexer) skipWhitespace() {
	for slices.Contains([]byte{' ', '\t', '\n', '\r'}, l.currChar) {
		l.nextChar()
	}
}

func (l *Lexer) nextToken() Token {
	l.skipWhitespace()

	var token Token
	switch l.currChar {
	case 0:
		token = Token{Type: EOFToken, Value: ""}
	case '+':
		token = Token{Type: PlusToken, Value: "+"}
	case '-':
		token = Token{Type: MinusToken, Value: "-"}
	case '*':
		token = Token{Type: AsteriskToken, Value: "*"}
	case '/':
		token = Token{Type: SlashToken, Value: "/"}
	case '(':
		token = Token{Type: LeftParenthesisToken, Value: "("}
	case ')':
		token = Token{Type: RightParenthesisToken, Value: ")"}
	default:
		if l.isDigit(l.currChar) {
			token = Token{Type: IntegerToken, Value: string(l.currChar)}
		} else {
			token = Token{Type: IllegalToken, Value: string(l.currChar)}
		}
	}

	l.nextChar()
	return token
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Lex() []Token {
	tokens := []Token{}
	for tok := l.nextToken(); tok.Type != EOFToken; tok = l.nextToken() {
		tokens = append(tokens, tok)
	}
	return tokens
}
