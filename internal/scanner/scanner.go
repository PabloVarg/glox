package scanner

import (
	"fmt"

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
	switch seen := s.advance(); seen {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	case '!':
		tokenType := token.BANG
		if s.match('=') {
			tokenType = token.BANG_EQUAL
		}

		s.addToken(tokenType, nil)
	case '<':
		tokenType := token.LESS
		if s.match('=') {
			tokenType = token.LESS_EQUAL
		}

		s.addToken(tokenType, nil)
	case '>':
		tokenType := token.GREATER
		if s.match('=') {
			tokenType = token.GREATER_EQUAL
		}

		s.addToken(tokenType, nil)
	case '/':
		if s.match('/') {
			s.advanceToNextLine()
			break
		}

		s.addToken(token.SLASH, nil)
	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		s.line++
		break
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

func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	s.Tokens = append(s.Tokens, token.Token{
		Type:    tokenType,
		Lexeme:  string(s.source[s.start:s.current]),
		Line:    s.line,
		Literal: literal,
	})
}
