package generate

import (
	"log"
	"strings"

	"github.com/krishna-godoi/gopher-maestro/ast"
	"github.com/krishna-godoi/gopher-maestro/token"
)

func GenerateForStatement(args, scope string) *ast.ForStatement {
	t := token.Token{
		Type:    token.FOR,
		Literal: "for",
	}

	forNode := ast.ForStatement{
		Token:   t,
		Context: "root",
	}

	parsedArgs := ParseArgs(args)
	if len(parsedArgs) != 3 {
		log.Fatal("Wrong number or args passed to FOR")
	}

	parsedScope := strings.Split(scope, ",")
	if len(parsedScope) > 0 {
		for i := range parsedScope {
			scopeItem := strings.TrimSpace(parsedScope[i])
			if len(scopeItem) > 0 {
				forNode.Scope = append(forNode.Scope, CallGenerator(scopeItem))
			}
		}
	}

	if len(parsedArgs[0]) > 0 {
		forNode.Variable = GenerateVarStatement(parsedArgs[0])
	}

	if len(parsedArgs[1]) > 0 {
		forNode.Condition = parsedArgs[1]
	}

	if len(parsedArgs[2]) > 0 {
		forNode.Increment = parsedArgs[1]
	}

	return &forNode
}
