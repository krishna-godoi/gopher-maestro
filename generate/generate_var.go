package generate

import (
	"log"

	"github.com/krishna-godoi/gopher-maestro/ast"
	"github.com/krishna-godoi/gopher-maestro/token"
)

func GenerateVarStatement(args string) *ast.VarStatement {
	tok := token.Token{
		Type:    token.VAR,
		Literal: "var",
	}

	parsedArgs := ParseArgs(args)

	if len(parsedArgs) != 3 {
		log.Fatal("Wrong number of arguments passed to VAR")
	}

	if len(parsedArgs[1]) == 0 && len(parsedArgs[2]) == 0 {
		log.Fatal("The type and value of a VAR can't both be empty")
	}

	varNode := ast.VarStatement{
		Token: tok,
		Type:  parsedArgs[1],
		Identifier: ast.Identifier{
			Token: token.IDENT,
		},
		Value: parsedArgs[2],
	}

	if len(parsedArgs[0]) > 0 {
		varNode.Identifier.Name = parsedArgs[0]
	} else {
		varNode.Identifier.Name = GenerateString()
	}

	return &varNode
}
