package generate

import (
	"github.com/krishna-godoi/gopher-ipsum/ast"
	"github.com/krishna-godoi/gopher-ipsum/token"
)

func GenerateFuncStatement() *ast.FuncStatement {
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

	for i := 0; i < 2; i++ {
		funcNode.Children = append(funcNode.Children, GenerateVarStatement())
	}

	return &funcNode
}
