package token

type Token struct {
	Type  string
	Literal string
}

const (
	FUNC    = "FUNC"
	VAR     = "VAR"
	IDENT   = "IDENT"
	TYPE    = "TYPE"
	NEWLINE = "\r\n"
	COMMA   = ","
	PLUS    = "+"
	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	ASSIGN  = "="
)
