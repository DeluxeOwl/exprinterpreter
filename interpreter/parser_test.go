package interpreter

import "testing"

func TestParserInitializesTokens(t *testing.T) {
	l := NewLexer("1 + 2")
	p := NewParser(l)

	if p.currToken.Type != IntegerToken && p.currToken.Value != "1" {
		t.Errorf("Expected integer token with value '1', got %+v", p.currToken)
	}

	if p.nextToken.Type != PlusToken && p.nextToken.Value != "+" {
		t.Errorf("Expected plus token with value '+', got %+v", p.nextToken)
	}
}

func TestParser(t *testing.T) {
	tests := []struct {
		source string
		parsed string
	}{
		{"1 + 2 * 3 + -123", "((1 + (2 * 3)) + (-123))"},
		{"-11 + 12", "((-11) + 12)"},
		{"", ""},
		{"---4", "(-(-(-4)))"},
	}

	for _, tt := range tests {
		l := NewLexer(tt.source)
		p := NewParser(l)
		program := p.ParseProgram().String()
		if program != tt.parsed {
			t.Errorf("Expected parsed expression '%s', got '%s'", tt.parsed, program)
		}
	}
}
