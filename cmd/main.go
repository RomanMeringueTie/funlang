package main

import (
	"fmt"
	"funlang/internal/parser"
)

func main() {
	fmt.Println("Hello from funclang")

	ast := parser.MockProgram()
	ast.Run()
}
