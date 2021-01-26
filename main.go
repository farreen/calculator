package main

import (
//	"fmt"
)

func main() {
	//calculateUsing3Operands()
	//calculateUsing3OperandsAndSameOp()
	//CalculateUsingSwitch(6, "+", 2)
	//calculateUsingIfelse()
	//CalculateUsingMap()
}
func CreateOperationMap() map[string]func(a, b int) int {
	m := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
		"%": func(a, b int) int { return a % b },
	}
	return m
}

func CalculateUsingMap(a int, op string, b int) int {
	m := CreateOperationMap()
	return m[op](a, b)
}

func CalculateUsingSwitch(a int, op string, b int) int {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		return -1
	}
}

func CalculateUsingIfelse(a int, op string, b int) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		return a / b
	} else if op == "%" {
		return a % b
	} else {
		return 0
	}
}

func CalculateUsing3OperandsAndSameOp(a int, op1 string, b int, op2 string, c int) int {
	m := CreateOperationMap()
	r := m[op1](a, b)
	return m[op2](r, c)
}

func CalculateUsing3Operands(a int, op1 string, b int, op2 string, c int) int {
	m := CreateOperationMap()
	if op1 == "/" || op1 == "*" {
		r := m[op1](a, b)
		return m[op2](r, c)
	} else if op2 == "*" || op2 == "/" {
		r := m[op2](b, c)
		return m[op1](a, r)
	}
	r := m[op1](a, b)
	return m[op2](r, c)
}
