package interpreter

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
