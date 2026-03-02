package main

import (
	"fmt"
	"foo/lexer"
	"foo/parser"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	content, _ := os.ReadFile("foo.fin")
	in := lexer.Tokenize(string(content))
	litter.Dump(in)
	fmt.Println(" ")
	fmt.Println("INICIANDO PARSER")
	p := parser.Parse(in)
	litter.Dump(p)
}
