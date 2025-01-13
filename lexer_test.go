package main

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		source string
		want   []Token
	}{
		{"1 + 2 * 3", []Token{
			{Type: IntegerToken, Value: "1"},
			{Type: PlusToken, Value: "+"},
			{Type: IntegerToken, Value: "2"},
			{Type: AsteriskToken, Value: "*"},
			{Type: IntegerToken, Value: "3"},
		}},
		{"-1-(3/   2)", []Token{
			{Type: MinusToken, Value: "-"},
			{Type: IntegerToken, Value: "1"},
			{Type: MinusToken, Value: "-"},
			{Type: LeftParenthesisToken, Value: "("},
			{Type: IntegerToken, Value: "3"},
			{Type: SlashToken, Value: "/"},
			{Type: IntegerToken, Value: "2"},
			{Type: RightParenthesisToken, Value: ")"},
		}},
	}

	for _, tt := range tests {
		l := NewLexer(tt.source)
		tokens := l.Lex()
		if !reflect.DeepEqual(tokens, tt.want) {
			t.Errorf("Lex(%q) = %v; want %v", tt.source, tokens, tt.want)
		}
	}
}
