package scanner

import "github.com/PabloVarg/glox/internal/token"

func ScanTokens(source string) []token.Token {
	tokens := []token.Token{}
	start, current, line := 0, 0, 1

	for {
		if isAtEnd(current, source) {
			break
		}

		start = current
		_ = start
		// scanToken()
	}

	tokens = append(tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Line:    line,
		Literal: nil,
	})
	return tokens
}

func isAtEnd(current int, source string) bool {
	return len(source) <= current
}
