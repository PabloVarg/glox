package scanner

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PabloVarg/glox/internal/reports"
	"github.com/PabloVarg/glox/internal/token"
)

type Scanner struct {
	Tokens  []token.Token
	source  string
	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens(source string) []token.Token {
	s.source = source
	s.Tokens = []token.Token{}
	s.start, s.current, s.line = 0, 0, 1

	for {
		if s.isAtEnd() {
			break
		}

		s.start = s.current
		_ = s.start
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Line:    s.line,
		Literal: nil,
	})
	return s.Tokens
}

func (s *Scanner) isAtEnd() bool {
	return len(s.source) <= s.current
}

func (s *Scanner) scanToken() {
	seen := s.advance()
	switch {
	case seen == '(':
		s.addToken(token.LEFT_PAREN, nil)
	case seen == ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case seen == '{':
		s.addToken(token.LEFT_BRACE, nil)
	case seen == '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case seen == ',':
		s.addToken(token.COMMA, nil)
	case seen == '.':
		s.addToken(token.DOT, nil)
	case seen == '-':
		s.addToken(token.MINUS, nil)
	case seen == '+':
		s.addToken(token.PLUS, nil)
	case seen == ';':
		s.addToken(token.SEMICOLON, nil)
	case seen == '*':
		s.addToken(token.STAR, nil)
	case seen == '!':
		tokenType := token.BANG
		if s.match('=') {
			tokenType = token.BANG_EQUAL
		}

		s.addToken(tokenType, nil)
	case seen == '<':
		tokenType := token.LESS
		if s.match('=') {
			tokenType = token.LESS_EQUAL
		}

		s.addToken(tokenType, nil)
	case seen == '>':
		tokenType := token.GREATER
		if s.match('=') {
			tokenType = token.GREATER_EQUAL
		}

		s.addToken(tokenType, nil)
	case seen == '/':
		if s.match('/') {
			s.advanceToNextLine()
			break
		}

		s.addToken(token.SLASH, nil)
	case seen == '"':
		s.string()
	case seen == ' ':
		break
	case seen == '\t':
		break
	case seen == '\r':
		break
	case seen == '\n':
		s.line++
		break
	case isDigit(seen):
		s.number()
	case isAlpha(seen):
		s.identifier()
	default:
		reports.Error(s.line, fmt.Sprintf("Unexpected character (%c).", seen))
	}
}

func (s *Scanner) advance() rune {
	result := rune(s.source[s.current])
	s.current++

	return result
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return rune(byte(0))
	}
	return rune(s.source[s.current])
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return rune(byte(0))
	}

	return rune(s.source[s.current+1])
}

func (s *Scanner) match(character rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != byte(character) {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advanceToNextLine() {
	for {
		if s.isAtEnd() || s.peek() == '\n' {
			break
		}

		s.advance()
	}
}

func (s *Scanner) string() {
	for {
		if s.peek() == '"' || s.isAtEnd() {
			break
		}

		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		reports.Error(s.line, "Unterminated string.")
		return
	}

	s.advance()
	s.addToken(token.STRING, string(s.source[s.start+1:s.current-1]))
}

func (s *Scanner) number() {
	log.Println("inn")
	for {
		if !isDigit(s.peek()) {
			break
		}

		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for {
			if !isDigit(s.peek()) {
				break
			}

			s.advance()
		}
	}

	literal, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		reports.Error(s.line, fmt.Sprintf("Could not parse %s", s.source[s.start:s.current]))
	}
	s.addToken(token.NUMBER, literal)
}

func (s *Scanner) identifier() {
	for {
		if !isAlphaNumeric(s.peek()) {
			break
		}

		s.advance()
	}

	seen := s.source[s.start:s.current]
	if tokenType, ok := token.KEYWORDS[seen]; ok {
		s.addToken(tokenType, nil)
		return
	}

	s.addToken(token.IDENTIFIER, nil)
}

func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	s.Tokens = append(s.Tokens, token.Token{
		Type:    tokenType,
		Lexeme:  string(s.source[s.start:s.current]),
		Line:    s.line,
		Literal: literal,
	})
}
