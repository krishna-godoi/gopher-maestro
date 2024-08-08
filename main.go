package main

import (
	"fmt"

	"github.com/krishna-godoi/gopher-maestro/generate"
)

func main() {
	genString := "FOR(VAR(i,,0),i<10, i++)[VAR(myVar,,15), VAR(myVar2,int,10)]"

	fmt.Println(generate.CallGenerator(genString))
}

