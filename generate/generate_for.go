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

	parsedScope := ParseArgs(scope)

	if len(parsedScope) > 0 {
		for i := range parsedScope {
			scopeItem := strings.TrimSpace(parsedScope[i])
			if len(scopeItem) > 0 {
				forNode.Scope = append(forNode.Scope, CallGenerator(scopeItem))
			}
		}
	}

	lVar, lCond, lInc := len(parsedArgs[0]), len(parsedArgs[1]), len(parsedArgs[2])

	if lVar > 0 || lInc > 0 {
		if lVar == 0 || lInc == 0 || lCond == 0 {
			log.Fatal("Illegal boomer loop setup")
		}
	}

	if lVar > 0 {
		_, args, _ := SplitGeneratorStatement(parsedArgs[0])
		forNode.Variable = GenerateVarStatement(args)
	}

	if lCond > 0 {
		forNode.Condition = parsedArgs[1]
	}

	if lInc > 0 {
		forNode.Increment = parsedArgs[2]
	}

	return &forNode
}
