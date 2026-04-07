package parser

import "fmt"

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
	id     string
	params []string
	expr   Expr
}

func (funDef FunDef) Run() error {
	return addFun(funDef.id, funDef.params, funDef.expr)
}

// Expressions
type Expr interface {
	Eval() uint
}

type Num struct {
	number uint
}

func (num Num) Eval() uint {
	return num.number
}

type Id struct {
	name string
}

func (id Id) Eval() uint {
	for _, variable := range *currentScope {
		if variable.name == id.name {
			return uint(variable.value)
		}
	}

	return 0
}

// Operations
type Plus struct {
	left  Expr
	right Expr
}

func (plus Plus) Eval() uint {
	return plus.left.Eval() + plus.right.Eval()
}

type Minus struct {
	left  Expr
	right Expr
}

func (minus Minus) Eval() uint {
	return minus.left.Eval() - minus.right.Eval()
}

type Div struct {
	left  Expr
	right Expr
}

func (div Div) Eval() uint {
	return div.left.Eval() / div.right.Eval()
}

type Mul struct {
	left  Expr
	right Expr
}

func (mul Mul) Eval() uint {
	return mul.left.Eval() * mul.right.Eval()
}

type Mod struct {
	left  Expr
	right Expr
}

func (mod Mod) Eval() uint {
	return mod.left.Eval() % mod.right.Eval()
}

type FunInv struct {
	id   string
	args []Expr
}

func (funInv FunInv) Eval() uint {
	mangledName := getFuncMangledName(funInv.id, uint(len(funInv.args)))
	currentScope = FunIds[mangledName].scope
	mapArgsToParams(mangledName, funInv.args)
	return FunIds[mangledName].expression.Eval()
}

func (funInv FunInv) Run() error {
	fmt.Println(funInv.Eval())
	return nil
}
