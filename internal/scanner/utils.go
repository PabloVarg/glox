package scanner

func isDigit(character rune) bool {
	return '0' <= character && character <= '9'
}

func isAlpha(character rune) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func isAlphaNumeric(character rune) bool {
	return isDigit(character) || isAlpha(character)
}
