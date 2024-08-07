package generate

import (
	"github.com/krishna-godoi/gopher-maestro/ast"
	"github.com/krishna-godoi/gopher-maestro/token"
)

func GenerateFuncStatement(args, scope string) *ast.FuncStatement {
	t := token.Token{
		Type:    token.FUNC,
		Literal: "func",
	}

	funcNode := ast.FuncStatement{
		Token:   t,
		Context: "root",
		Name: ast.Identifier{
			Token: token.IDENT,
			Name:  GenerateString(),
		},
	}

	return &funcNode
}
