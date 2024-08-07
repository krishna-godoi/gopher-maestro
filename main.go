package main

import (
	"fmt"

	"github.com/krishna-godoi/gopher-ipsum/generate"
)

func main() {
	p := generate.GenerateRoot()
	for i := range p.Statements {
		fmt.Println(p.Statements[i])
	}
}
