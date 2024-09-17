package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPARENT = "("
	RPARENT = ")"
	LBRACE  = "{"
	RBARACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
