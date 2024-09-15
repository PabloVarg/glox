package token

import "fmt"

type TokenType string

const (
	// Single character tokens
	LEFT_PAREN  = TokenType("left_paren")
	RIGHT_PAREN = TokenType("right_paren")
	LEFT_BRACE  = TokenType("left_brace")
	RIGHT_BRACE = TokenType("right_brace")
	COMMA       = TokenType("comma")
	DOT         = TokenType("dot")
	MINUS       = TokenType("minus")
	PLUS        = TokenType("plus")
	SEMICOLON   = TokenType("semicolon")
	SLASH       = TokenType("slash")
	STAR        = TokenType("star")

	// One or two character tokens
	BANG          = TokenType("bang")
	BANG_EQUAL    = TokenType("bang_equal")
	EQUAL         = TokenType("equal")
	EQUAL_EQUAL   = TokenType("equal_equal")
	GREATER       = TokenType("greater")
	GREATER_EQUAL = TokenType("greater_equal")
	LESS          = TokenType("less")
	LESS_EQUAL    = TokenType("less_equal")

	// Literals
	IDENTIFIER = TokenType("identifier")
	STRING     = TokenType("string")
	NUMBER     = TokenType("number")

	// Keywords
	AND    = TokenType("and")
	CLASS  = TokenType("class")
	ELSE   = TokenType("else")
	FALSE  = TokenType("false")
	FUN    = TokenType("fun")
	FOR    = TokenType("for")
	IF     = TokenType("if")
	NIL    = TokenType("nil")
	OR     = TokenType("or")
	PRINT  = TokenType("print")
	RETURN = TokenType("return")
	SUPER  = TokenType("super")
	THIS   = TokenType("this")
	TRUE   = TokenType("true")
	VAR    = TokenType("var")
	WHILE  = TokenType("while")

	EOF = TokenType("EOF")
)

var KEYWORDS = map[string]TokenType{
	"and":    TokenType("and"),
	"class":  TokenType("class"),
	"else":   TokenType("else"),
	"false":  TokenType("false"),
	"fun":    TokenType("fun"),
	"for":    TokenType("for"),
	"if":     TokenType("if"),
	"nil":    TokenType("nil"),
	"or":     TokenType("or"),
	"print":  TokenType("print"),
	"return": TokenType("return"),
	"super":  TokenType("super"),
	"this":   TokenType("this"),
	"true":   TokenType("true"),
	"var":    TokenType("var"),
	"while":  TokenType("while"),
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Line    int
	Literal any
}

func (token Token) toString() string {
	return fmt.Sprintf("%s, %s, %v", token.Type, token.Lexeme, token.Literal)
}
