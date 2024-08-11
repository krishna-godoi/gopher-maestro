package token

type Token struct {
	Type    string
	Literal string
}

const (
	FUNC    = "FUNC"
	FOR     = "FOR"
	VAR     = "VAR"
	IF      = "IF"
	ELSE    = "ELSE"
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
