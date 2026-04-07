package parser

import (
	"fmt"
	"slices"
)

type Scope map[string]int

type funMetadata struct {
	expression Expr
	scope      *Scope
}

var FunIds map[string]funMetadata = make(map[string]funMetadata)

func (scope *Scope) addVarId(id string) error {
	if _, isPresent := (*scope)[id]; isPresent {
		return fmt.Errorf("variable with identifier %s already exists in this scope", id)
	}

	(*scope)[id] = -1
	return nil
}

func (scope *Scope) setVarValue(id string, value uint) error {
	if _, isPresent := (*scope)[id]; !isPresent {
		return fmt.Errorf("variable with identifier %s not exists in this scope", id)
	}

	if value >= 0 {
		(*scope)[id] = int(value)
	}
	return nil
}

func addFun(id string, params []string, expr Expr, scope *Scope) error {
	if _, isPresent := FunIds[id]; isPresent {
		return fmt.Errorf("variable with identifier %s already exists in this scope", id)
	}

	validateExpr(id, params, expr)

	mangledName := getFuncMangledName(id, uint(len(params)))
	FunIds[mangledName] = funMetadata{expr, scope}

	return nil
}

func getFuncMangledName(id string, argCount uint) string {
	return fmt.Sprintf("%s%d", id, argCount)
}

func validateExpr(funId string, params []string, expr Expr) error {
	id, ok := expr.(Id)
	if ok {
		if !slices.Contains(params, id.name) {
			return fmt.Errorf("var %s in expression of function %s is undeclared", id, funId)
		}
	}

	//: Add processing of operations

	return nil
}

// ===DEBUG===

func printFunMap() {
	for key := range FunIds {
		fmt.Printf("key: %s\n", key)
	}
}
