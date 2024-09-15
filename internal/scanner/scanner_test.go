package scanner

import (
	"testing"

	"github.com/PabloVarg/glox/internal/token"
)

func RecognizeSingleToken(t *testing.T, source, lexeme string, tokenType token.TokenType, literal any) {
	var scan Scanner
	scan.ScanTokens(source)

	if len(scan.Tokens) != 2 {
		t.Errorf("Wrong number of tokens, encountered %d, expected %d", len(scan.Tokens), 2)
	}

	if scan.Tokens[0].Type != tokenType {
		t.Errorf("Wrong token type, encountered %s, expected %s", scan.Tokens[0].Type, tokenType)
	}

	if scan.Tokens[0].Lexeme != lexeme {
		t.Errorf("Wrong lexeme, detected %s, expected %s", scan.Tokens[0].Lexeme, lexeme)
	}

	if scan.Tokens[0].Literal != literal {
		t.Errorf("Wrong literal, encountered %v, expected %v", scan.Tokens[0].Literal, literal)
	}
}

func TestSingleCharacters(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
		}
	}{
		{
			source: ".",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    ".",
				tokenType: token.DOT,
			},
		},
		{
			source: "{",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "{",
				tokenType: token.LEFT_BRACE,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, nil)
		})
	}
}

func TestTwoCharacters(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
		}
	}{
		{
			source: "!=",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "!=",
				tokenType: token.BANG_EQUAL,
			},
		},
		{
			source: "<=",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "<=",
				tokenType: token.LESS_EQUAL,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, nil)
		})
	}
}

func TestIdentifiers(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
		}
	}{
		{
			source: "asontheu",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "asontheu",
				tokenType: token.IDENTIFIER,
			},
		},
		{
			source: "_abc",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "_abc",
				tokenType: token.IDENTIFIER,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, nil)
		})
	}
}

func TestKeywords(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
		}
	}{
		{
			source: "for",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "for",
				tokenType: token.FOR,
			},
		},
		{
			source: "if",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
			}{
				lexeme:    "if",
				tokenType: token.IF,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, nil)
		})
	}
}

func TestStrings(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
			literal   string
		}
	}{
		{
			source: "\"test\"",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
				literal   string
			}{
				lexeme:    "\"test\"",
				tokenType: token.STRING,
				literal:   "test",
			},
		},
		{
			source: "\"abc\"",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
				literal   string
			}{
				lexeme:    "\"abc\"",
				tokenType: token.STRING,
				literal:   "abc",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, tC.expected.literal)
		})
	}
}

func TestNumbers(t *testing.T) {
	testCases := []struct {
		source   string
		expected struct {
			lexeme    string
			tokenType token.TokenType
			literal   float64
		}
	}{
		{
			source: "25.2",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
				literal   float64
			}{
				lexeme:    "25.2",
				tokenType: token.NUMBER,
				literal:   25.2,
			},
		},
		{
			source: "0.0",
			expected: struct {
				lexeme    string
				tokenType token.TokenType
				literal   float64
			}{
				lexeme:    "0.0",
				tokenType: token.NUMBER,
				literal:   0,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.source, func(t *testing.T) {
			RecognizeSingleToken(t, tC.source, tC.expected.lexeme, tC.expected.tokenType, tC.expected.literal)
		})
	}
}
