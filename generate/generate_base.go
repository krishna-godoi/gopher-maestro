package generate

import (
	"bufio"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/krishna-godoi/gopher-maestro/ast"
)

func CallGenerator(str string) ast.Statement {
	var stmt ast.Statement
	var genKey, args, scope string

	argsStart := strings.Index(str, "(")
	scopeStart := strings.Index(str, "[")

	if argsStart == -1 && scopeStart == -1 {
		genKey = str
	} else if argsStart == -1 {
		genKey = str[:scopeStart]
		scope = str[scopeStart+1 : len(str)-1]
	} else {
		genKey = str[:argsStart]

		argsEnd := strings.Index(str, ")")
		args = str[argsStart+1 : argsEnd]

		if scopeStart != -1 {
			scope = str[scopeStart+1 : len(str)-1]
		}
	}

	switch genKey {
	case "VAR":
		stmt = GenerateVarStatement(args)
	case "FUNC":
		stmt = GenerateFuncStatement(args, scope)
	case "FOR":
		stmt = GenerateForStatement(args, scope)
	}

	return stmt
}

func ParseArgs(args string) []string {
	parsedArgs := strings.Split(args, ",")

	for i := range parsedArgs {
		parsedArgs[i] = strings.TrimSpace(parsedArgs[i])
	}

	return parsedArgs
}

func GenerateString() string {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	words := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	idx := rand.IntN(len(words))

	return words[idx]
}

var operators = map[int]string{
	0: "+",
	1: "-",
	2: "*",
	3: "/",
	4: "%",
}

var possibilities = map[int]string{
	0: "INT",
	1: "OP",
}

func GenerateMathExpr(depth int) string {
	// TODO: Extract the decay rate to configs
	decay := float64(depth) * 0.2
	seed := (rand.Float64() - decay) * math.Floor(float64(len(possibilities)))
	choice := possibilities[int(seed)]

	if choice == "INT" {
		return GenerateInt()
	}

	op := rand.IntN(len(operators))
	expr := GenerateMathExpr(depth+1) + operators[op] + GenerateMathExpr(depth+1)

	parentheses := rand.Float64()
	// TODO: Separate parentheses decay rate from the root one
	if parentheses > 1-decay {
		return "(" + expr + ")"
	}

	return expr
}

func GenerateInt() string {
	return strconv.Itoa(rand.IntN(101))
}
