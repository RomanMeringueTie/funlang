package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func parseFile(path string) AST {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokens := make(AST, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens = append(tokens, parseLine(line)...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}

func parseLine(line string) AST {
	if len(strings.TrimSpace(line)) == 0 {
		return AST{}
	}

	if strings.Contains(line, "=") {
		return parseFunDefinition(line)
	} else {
		return parseFunInvocation(line)
	}

}

func parseFunDefinition(line string) AST {
	//: Implement parseFunDefinition
	var funcIdIndex int
	for index, char := range line {
		if char == '(' {
			funcIdIndex = index
		}
	}

	funcIdIndex++
	return AST{}
}

func parseFunInvocation(line string) AST {
	//: Implement parseFunInvocation
	var funcIdIndex int
	for index, char := range line {
		if char == '(' {
			funcIdIndex = index
		}
	}

	funcIdIndex++

	return AST{}
}

func MockProgram() AST {
	// f() = 1
	f := FunDef{id: "f", params: []string{}, expr: Num{1}}
	// f(a) = 2
	f1 := FunDef{id: "f", params: []string{"a"}, expr: Id{"a"}}
	// f(a, b) = a + b
	f2 := FunDef{id: "f", params: []string{"a", "b"}, expr: Plus{left: Id{"a"}, right: Id{"b"}}}
	// f()
	fCall := FunInv{id: "f", args: []Expr{}}
	// f(1)
	f1Call := FunInv{id: "f", args: []Expr{Num{1}}}
	// f(10, 59)
	f2Call := FunInv{id: "f", args: []Expr{Num{11}, Num{2}}}

	ast := AST{f, f1, f2, fCall, f1Call, f2Call}
	return ast
}
