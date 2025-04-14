package interpreter

import (
	"fmt"
	"go-ts-compiler/parser"
)

type Environment struct{
	store map[string]int
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]int),
	}
}

func (e *Environment) set(name string, value int){
	e.store[name]= value
}

func (e *Environment) get(name string)( int, bool){
	val, ok := e.store[name]
	return val, ok
}

func Eval(node parser.Node, env *Environment) int{
	switch n := node.(type){
	case *parser.StatementNode:
		val := Eval(n.Value, env)
		env.set(n.Variable, val)
		return val
	case *parser.ExpressionNode:
		left := atoi(n.Left.String())
		right := atoi(n.Right.String())
		switch n.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
	}
 }
 return 0
}

func atoi(lit string) int {
	var val int
	fmt.Sscanf(lit, "%d", &val)
	return val
}