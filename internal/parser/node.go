package parser

import (
	"fmt"
	"log"
)

var currentScope *Scope = nil

type AST []Node

func (ast *AST) Run() error {
	for _, node := range *ast {
		if err := node.Run(); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

type Node interface {
	Run() error
}

type FunDef struct {
	Id     string
	Params []string
	Expr   Expr
}

func (funDef FunDef) Run() error {
	return addFun(funDef.Id, funDef.Params, funDef.Expr)
}

// Expressions
type Expr interface {
	Eval() uint
}

type Num struct {
	Number uint
}

func (num Num) Eval() uint {
	return num.Number
}

type Id struct {
	Name string
}

func (id Id) Eval() uint {
	value := -1
	for _, variable := range *currentScope {
		if variable.name == id.Name {
			value = variable.value
			break
		}
	}

	if value < 0 {
		log.Fatalf("parameter %s is not initialized", id.Name)
	}

	return uint(value)
}

// Operations
type Plus struct {
	Left  Expr
	Right Expr
}

func (plus Plus) Eval() uint {
	return plus.Left.Eval() + plus.Right.Eval()
}

type Minus struct {
	Left  Expr
	Right Expr
}

func (minus Minus) Eval() uint {
	return minus.Left.Eval() - minus.Right.Eval()
}

type Div struct {
	Left  Expr
	Right Expr
}

func (div Div) Eval() uint {
	return div.Left.Eval() / div.Right.Eval()
}

type Mul struct {
	Left  Expr
	Right Expr
}

func (mul Mul) Eval() uint {
	return mul.Left.Eval() * mul.Right.Eval()
}

type Mod struct {
	Left  Expr
	Right Expr
}

func (mod Mod) Eval() uint {
	return mod.Left.Eval() % mod.Right.Eval()
}

type FunInv struct {
	Id   string
	Args []Expr
}

func (funInv FunInv) Eval() uint {
	mangledName := getFuncMangledName(funInv.Id, uint(len(funInv.Args)))
	currentScope = FunIds[mangledName].scope
	mapArgsToParams(mangledName, funInv.Args)
	return FunIds[mangledName].expression.Eval()
}

func (funInv FunInv) Run() error {
	fmt.Println(funInv.Eval())
	return nil
}
