package generate

import (
	"bufio"
	"math"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/krishna-godoi/gopher-ipsum/ast"
	"github.com/krishna-godoi/gopher-ipsum/token"
)

type PNode struct {
	Token  string
	Weight float64
}

func GenerateRoot() ast.Program {
	l := 2
	possible := []PNode{
		{Token: token.FUNC, Weight: 1.00},
		{Token: token.VAR, Weight: 0.35},
	}

	tree := ast.Program{}

	for l > 0 {
		idx := rand.IntN(l)
		chance := rand.Float64()

		possible[idx].Weight = max(0, possible[idx].Weight-chance)

		if possible[idx].Weight == 0 {
			possible[idx], possible[l-1] = possible[l-1], possible[idx]
			l = l - 1
		} else {
			tree.Statements = append(tree.Statements, GenerateNode(possible[idx].Token))
		}
	}

	return tree
}

func GenerateNode(t string) ast.Statement {
	if t == token.FUNC {
		return GenerateFuncStatement()
	} else {
		return GenerateVarStatement()
	}
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
