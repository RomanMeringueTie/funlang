package interpreter

import (
	"bufio"
	"fmt"
	"funlang/pkg/data_structures"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(path string) AST {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokens := make(AST, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		node := parseLine(line)
		if node != nil {
			tokens = append(tokens, node)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}

func parseLine(line string) Node {
	if len(strings.TrimSpace(line)) == 0 {
		return nil
	}

	if strings.Contains(line, "=") {
		return parseFunDefinition(line)
	} else {
		return parseFunInvocation(line)
	}

}

func parseFunDefinition(line string) FunDef {
	var id string
	var params []string
	var expr Expr

	var leftParIndex = 0
	for leftParIndex := range line {
		if line[leftParIndex] == '(' {
			id = line[:leftParIndex]
		}
	}

	var rightParIndex = 0
	for {
		if line[rightParIndex] == ')' {
			paramsString := line[leftParIndex+2 : rightParIndex]
			params = getParams(paramsString)
			break
		}
		rightParIndex++
	}

	var equalsIndex = rightParIndex + 1
	for {
		if line[equalsIndex] == '=' {
			break
		}

		equalsIndex++
	}

	expr = parsePolishNotation(line[equalsIndex+1:])

	funDef := FunDef{Id: id, Params: params, Expr: expr}
	fmt.Println(funDef)

	return funDef
}

func getParams(paramsString string) []string {
	paramsStringWithoutSpaces := strings.ReplaceAll(paramsString, " ", "")

	params := strings.Split(paramsStringWithoutSpaces, ",")
	return params
}

func parsePolishNotation(expression string) Expr {
	expressionWithSingleSpaces := strings.TrimSpace(strings.Join(strings.Fields(expression), " "))
	splittedExpression := strings.Split(expressionWithSingleSpaces, " ")
	operationStack := data_structures.NewStack[string]()
	variableStack := data_structures.NewStack[string]()

	for _, token := range splittedExpression {
		if isOperation(token) {
			operationStack.Push(token)
		} else {
			variableStack.Push(token)
		}
	}

	return getExpression(operationStack, variableStack)
}

func getExpression(operationStack, variableStack *data_structures.Stack[string]) Expr {

	for {
		if operationStack.Size() == 0 {
			return idOrNumberToExpr(variableStack.Pop())
		}

		operation := operationStack.Pop()
		right := idOrNumberToExpr(variableStack.Pop())

		switch operation {
		case "+":
			return Plus{Left: getExpression(operationStack, variableStack), Right: right}
		case "-":
			return Minus{Left: getExpression(operationStack, variableStack), Right: right}
		case "*":
			return Mul{Left: getExpression(operationStack, variableStack), Right: right}
		case "/":
			return Div{Left: getExpression(operationStack, variableStack), Right: right}
		case "%":
			return Mod{Left: getExpression(operationStack, variableStack), Right: right}
		}
	}

}

func isOperation(token string) bool {
	switch token {
	case "+", "-", "*", "/", "%":
		return true
	default:
		return false
	}
}

func idOrNumberToExpr(idOrNumber string) Expr {
	value, err := strconv.Atoi(idOrNumber)
	if err != nil {
		return Id{Name: idOrNumber}
	}

	return Num{Number: uint(value)}
}

func parseFunInvocation(line string) FunInv {
	var id string
	var args []Expr

	var leftParIndex = 0
	for leftParIndex := range line {
		if line[leftParIndex] == '(' {
			id = line[:leftParIndex]
		}
	}

	var rightParIndex = 0
	for {
		if line[rightParIndex] == ')' {
			argsString := line[leftParIndex+2 : rightParIndex]
			args = getArgs(argsString)
			break
		}
		rightParIndex++
	}
	funInv := FunInv{Id: id, Args: args}
	fmt.Println(funInv)

	return funInv
}

func getArgs(args string) []Expr {

	expessions := make([]Expr, 0)

	splittedArgs := strings.Split(args, ",")
	for _, arg := range splittedArgs {
		expessions = append(expessions, parsePolishNotation(arg))
	}

	return expessions
}
