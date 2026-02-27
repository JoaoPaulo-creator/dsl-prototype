package main

import (
	"fmt"
	"foo/lexer"
	"os"
)

func main() {
	content, _ := os.ReadFile("foo.fin")
	in := lexer.Tokenize(string(content))
	fmt.Println(in)
}
