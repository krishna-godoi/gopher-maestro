package generate

import (
	"math/rand/v2"

	"github.com/krishna-godoi/gopher-ipsum/ast"
	"github.com/krishna-godoi/gopher-ipsum/token"
)

func GenerateVarStatement() *ast.VarStatement {
	t := token.Token{
		Type:    token.VAR,
		Literal: "var",
	}

	varNode := ast.VarStatement{
		Token:   t,
		Context: "root",
		Name: ast.Identifier{
			Token: token.IDENT,
			Name:  GenerateString(),
		},
	}

	i := rand.IntN(2)

	if i == 0 {
		varNode.Value = GenerateString()
	} else {
		varNode.Value = GenerateMathExpr(0)
	}

	return &varNode
}
