package interpreter

import "bytes"

type Expression interface {
	String() string
}

type Program struct {
	Expressions []Expression
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, expr := range p.Expressions {
		out.WriteString(expr.String())
	}

	return out.String()
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type IntegerExpression struct {
	Token Token
	Value int64
}

func (il *IntegerExpression) String() string { return il.Token.Value }
