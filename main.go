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

func CalculateUsingMap(a int, op string, b int) int {
	m := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
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
	switch op1 {
	case "+":
		return a + b + c
	case "-":
		return a - b - c
	case "*":
		return a * b * c
	case "/":
		return a / b / c
	case "%":
		return a % b % c
	default:
		return -1
	}
}

func CalculateUsing3Operands(a int, op1 string, b int, op2 string, c int) int {
	if op1 == "+" && op2 == "+" {
		return a + b + c
	} else if op1 == "+" && op2 == "-" {
		return a + b - c
	} else if op1 == "+" && op2 == "*" {
		return a + b*c
	} else if op1 == "+" && op2 == "/" {
		return a + b/c
	} else if op1 == "-" && op2 == "+" {
		return a - b + c
	} else if op1 == "-" && op2 == "-" {
		return a - b - c
	} else if op1 == "-" && op2 == "*" {
		return a - b*c
	} else if op1 == "-" && op2 == "/" {
		return a - b/c
	} else if op1 == "*" && op2 == "+" {
		return a*b + c
	} else if op1 == "*" && op2 == "-" {
		return a*b - c
	} else if op1 == "*" && op2 == "*" {
		return a * b * c
	} else if op1 == "*" && op2 == "/" {
		return a * b / c
	} else if op1 == "/" && op2 == "+" {
		return a/b + c
	} else if op1 == "/" && op2 == "-" {
		return a/b - c
	} else if op1 == "/" && op2 == "*" {
		return a / b * c
	} else if op1 == "/" && op2 == "/" {
		return a / b / c
	} else {
		return 0
	}
}
