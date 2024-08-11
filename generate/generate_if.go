package generate

import (
	"github.com/krishna-godoi/gopher-maestro/ast"
	"github.com/krishna-godoi/gopher-maestro/token"
)

func GenerateIfStatement(args, scope string) *ast.IfStatement {
	tok := token.Token{
		Type:    token.IF,
		Literal: "if",
	}

	parsedArgs := ParseArgs(args)
	parsedScope := ParseArgs(scope)

	variable, condition, alternatives := parsedArgs[0], parsedArgs[1], parsedArgs[2]

	ifNode := ast.IfStatement{
		Token:     tok,
		Condition: condition,
	}

	if len(variable) > 0 {
		ifNode.Variable = CallGenerator(variable)
	}

	if len(alternatives) > 0 {
		var parsedAlts []string
		if alternatives[0] == '[' {
			parsedAlts = ParseArgs(alternatives[1 : len(alternatives)-1])
		} else {
			parsedAlts = ParseArgs(alternatives)
		}

		for i := range parsedAlts {
			alt := CallGenerator(parsedAlts[i])
			ifNode.Alternatives = append(ifNode.Alternatives, alt)
		}
	}

	if len(parsedScope) > 0 {
		for i := range parsedScope {
			ifNode.Scope = append(ifNode.Scope, CallGenerator(parsedScope[i]))
		}
	}

	return &ifNode
}
