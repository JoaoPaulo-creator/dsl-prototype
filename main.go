package main

import (
	"foo/lexer"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	content, _ := os.ReadFile("foo.fin")
	in := lexer.Tokenize(string(content))
	litter.Dump(in)
}
