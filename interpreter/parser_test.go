package interpreter

import "testing"

func TestParserConstructor(t *testing.T) {
	l := NewLexer("1 + 2")
	p := NewParser(l)

	if p.currToken.Type != IntegerToken && p.currToken.Value != "1" {
		t.Errorf("Expected integer token with value '1', got %+v", p.currToken)
	}

	if p.nextToken.Type != PlusToken && p.nextToken.Value != "+" {
		t.Errorf("Expected plus token with value '+', got %+v", p.nextToken)
	}
}
