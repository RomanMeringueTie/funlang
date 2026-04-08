package interpreter

import (
	"fmt"
	"log"
	"slices"
)

// Scope
type Var struct {
	name  string
	value int
}

type Scope []Var

func (scope *Scope) addVar(index uint, name string) {
	(*scope)[index] = Var{name: name, value: -1}
}

func (scope *Scope) setVarValue(index uint, value uint) {
	(*scope)[index].value = int(value)
}

type funData struct {
	expression Expr
	scope      *Scope
}

var currentScope *Scope = nil

// Symbols table
var FunIds map[string]funData = make(map[string]funData)

func addFun(id string, params []string, expr Expr) error {
	if _, isPresent := FunIds[id]; isPresent {
		return fmt.Errorf("function %s already exists", id)
	}

	mangledName := getFuncMangledName(id, uint(len(params)))
	scope := make(Scope, len(params))
	FunIds[mangledName] = funData{expr, &scope}

	err := validateFunParams(mangledName, params)
	if err != nil {
		return err
	}

	err = validateFunExpr(mangledName, params, expr)
	if err != nil {
		return err
	}

	return nil
}

// Checks
func validateFunParams(funId string, params []string) error {
	seen := make(map[string]struct{}, len(params))
	for index, param := range params {
		if _, ok := seen[param]; ok {
			return fmt.Errorf("parameter %s already exists in function %s", param, funId)
		}
		seen[param] = struct{}{}
		FunIds[funId].scope.addVar(uint(index), param)
	}

	return nil
}

func validateFunExpr(funId string, params []string, expr Expr) error {
	switch exprType := expr.(type) {
	case Id:
		if !slices.Contains(params, exprType.Name) {
			return fmt.Errorf("parameter %s in function %s is undeclared", exprType.Name, funId)
		}
		return nil
	case Num:
		return nil
	case Plus:
		err := validateFunExpr(funId, params, exprType.Left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.Right)
		if err != nil {
			return err
		}
	case Minus:
		err := validateFunExpr(funId, params, exprType.Left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.Right)
		if err != nil {
			return err
		}
	case Mul:
		err := validateFunExpr(funId, params, exprType.Left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.Right)
		if err != nil {
			return err
		}
	case Div:
		err := validateFunExpr(funId, params, exprType.Left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.Right)
		if err != nil {
			return err
		}
	case Mod:
		err := validateFunExpr(funId, params, exprType.Left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.Right)
		if err != nil {
			return err
		}
	}

	return nil
}

// Helpers
func mapArgsToParams(id string, args []Expr) {
	if _, isPresent := FunIds[id]; !isPresent {
		log.Fatalf("function %s not exists", id)
	}

	for index, expr := range args {
		FunIds[id].scope.setVarValue(uint(index), expr.Eval())
	}
}

func getFuncMangledName(id string, argCount uint) string {
	return fmt.Sprintf("%s%d", id, argCount)
}

// ===DEBUG===
func printFunMap() {
	for key := range FunIds {
		fmt.Printf("key: %s\n", key)
	}
}
