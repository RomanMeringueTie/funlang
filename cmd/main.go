package main

import (
	"fmt"
	"funlang/internal/interpreter"
)

func mockProgram() interpreter.AST {
	// f() = 1
	f := interpreter.FunDef{Id: "f", Params: []string{}, Expr: interpreter.Num{1}}
	// f(a) = 2
	f1 := interpreter.FunDef{Id: "f", Params: []string{"a"}, Expr: interpreter.Id{"a"}}
	// f(a, b) = a + b
	f2 := interpreter.FunDef{Id: "f", Params: []string{"a", "b"}, Expr: interpreter.Plus{Left: interpreter.Id{"a"}, Right: interpreter.Id{"b"}}}
	// f()
	fCall := interpreter.FunInv{Id: "f", Args: []interpreter.Expr{}}
	// f(1)
	f1Call := interpreter.FunInv{Id: "f", Args: []interpreter.Expr{interpreter.Num{1}, interpreter.Num{2}}}
	// f(10, 59)
	f2Call := interpreter.FunInv{Id: "f", Args: []interpreter.Expr{interpreter.Num{1}, interpreter.FunInv{Id: "f", Args: []interpreter.Expr{interpreter.Num{1}, interpreter.Num{2}}}}}

	ast := interpreter.AST{f, f1, f2, fCall, f1Call, f2Call}
	return ast
}

func main() {
	fmt.Println("Hello from funclang")

	ast := mockProgram()
	ast.Run()
}
