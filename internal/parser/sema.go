package parser

import (
	"fmt"
	"log"
	"slices"
)

type Var struct {
	string
	uint
}
type Scope []uint

type funMetadata struct {
	expression Expr
	scope      *Scope
}

var FunIds map[string]funMetadata = make(map[string]funMetadata)

func (scope *Scope) setVarValue(index int, value uint) {
	(*scope)[index] = value
}

func addFun(id string, params []string, expr Expr) error {
	if _, isPresent := FunIds[id]; isPresent {
		return fmt.Errorf("function %s already exists", id)
	}

	err := validateFunParams(id, params)
	if err != nil {
		return err
	}

	err = validateFunExpr(id, params, expr)
	if err != nil {
		return err
	}

	mangledName := getFuncMangledName(id, uint(len(params)))
	scope := make(Scope, len(params))
	FunIds[mangledName] = funMetadata{expr, &scope}

	return nil
}

func mapArgsToParams(id string, args []Expr) {
	if _, isPresent := FunIds[id]; !isPresent {
		log.Fatalf("function %s not exists", id)
	}

	for index, expr := range args {
		FunIds[id].scope.setVarValue(index, expr.Eval())
	}
}

func getFuncMangledName(id string, argCount uint) string {
	return fmt.Sprintf("%s%d", id, argCount)
}

func validateFunParams(funId string, params []string) error {
	seen := make(map[string]struct{}, len(params))
	for _, param := range params {
		if _, ok := seen[param]; ok {
			return fmt.Errorf("parameter %s already exists in function %s", param, funId)
		}
		seen[param] = struct{}{}
	}

	return nil
}

func validateFunExpr(funId string, params []string, expr Expr) error {
	switch exprType := expr.(type) {
	case Id:
		if !slices.Contains(params, exprType.name) {
			return fmt.Errorf("parameter %s in function %s is undeclared", exprType.name, funId)
		}
		return nil
	case Num:
		return nil
	case Plus:
		err := validateFunExpr(funId, params, exprType.left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.right)
		if err != nil {
			return err
		}
	case Minus:
		err := validateFunExpr(funId, params, exprType.left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.right)
		if err != nil {
			return err
		}
	case Mul:
		err := validateFunExpr(funId, params, exprType.left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.right)
		if err != nil {
			return err
		}
	case Div:
		err := validateFunExpr(funId, params, exprType.left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.right)
		if err != nil {
			return err
		}
	case Mod:
		err := validateFunExpr(funId, params, exprType.left)
		if err != nil {
			return err
		}
		err = validateFunExpr(funId, params, exprType.right)
		if err != nil {
			return err
		}
	}

	return nil
}

// ===DEBUG===

func printFunMap() {
	for key := range FunIds {
		fmt.Printf("key: %s\n", key)
	}
}
