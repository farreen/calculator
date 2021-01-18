package main

import (
	"fmt"
)

func main() {
	//calculateUsing3Operands()
	//calculateUsing3OperandsAndSameOp()
	//CalculateUsingSwitch(6, "+", 2)
	//calculateUsingIfelse()
	//fmt.Printf("sum = %d\n", AddTwoNum(2, 2))
	//CalculateUsingMap()
}

func CalculateUsingMap(a int, op string, b int) int {
	m := map[string]func(a, b int) int{
		"+": add,
		"-": minus,
		"*": multiply,
		"/": divide,
	}
	return m[op](a, b)
}

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func AddTwoNum(x, y int) int {
	return x + y
}

func CalculateUsingSwitch(a int, op string, b int) int {
	//var a, b int
	//var op string
	//	fmt.Print("enter two numbers\n")
	//	fmt.Scanf("%d %s %d", &a, &op, &b)
	switch op {
	case "+":
		//fmt.Println(a + b)
		return a + b
	case "-":
		//fmt.Println(a - b)
		return a - b
	case "*":
		//fmt.Println(a * b)
		return a * b
	case "/":
		//fmt.Println(a / b)
		return a / b
	case "%":
		//fmt.Println(a % b)
		return a % b
	default:
		//fmt.Println("not found")
		return -1
	}
}

func CalculateUsingIfelse(a int, op string, b int) int {
	//var a, b int
	//var op string
	//fmt.Print("enter two numbers\n")
	//fmt.Scanf("%d %s %d", &a, &op, &b)
	if op == "+" {
		//fmt.Println(a + b)
		return a + b
	} else if op == "-" {
		//	fmt.Println(a - b)
		return a - b
	} else if op == "*" {
		//	fmt.Println(a * b)
		return a * b
	} else if op == "/" {
		//	fmt.Println(a / b)
		return a / b
	} else if op == "%" {
		//	fmt.Println(a % b)
		return a % b
	} else {
		//	fmt.Println("not found")
		return 0
	}
}

func CalculateUsing3OperandsAndSameOp(a int, op1 string, b int, op2 string, c int) int {
	//var a, b, c int
	//var op1, op2 string
	//fmt.Print("enter three numbers\n")
	//fmt.Scanf("%d %s %d %s %d", &a, &op1, &b, &op2, &c)
	switch op1 {
	case "+":
		//	fmt.Println(a + b + c)
		return a + b + c
	case "-":
		//	fmt.Println(a - b - c)
		return a - b - c
	case "*":
		//	fmt.Println(a * b * c)
		return a * b * c
	case "/":
		//	fmt.Println(a / b / c)
		return a / b / c
	case "%":
		//	fmt.Println(a % b % c)
		return a % b % c
	default:
		// fmt.Println("not found")
		return -1
	}
}

func CalculateUsing3Operands(a int, op1 string, b int, op2 string, c int) int {
	//var a, b, c int
	//var op1, op2 string
	//fmt.Println("enter three numbers")
	//fmt.Scanf("%d %s %d %s %d", &a, &op1, &b, &op2, &c)
	if op1 == "+" && op2 == "+" {
		//	fmt.Println(a + b + c)
		return a + b + c
	} else if op1 == "+" && op2 == "-" {
		//	fmt.Println(a + b - c)
		return a + b - c
	} else if op1 == "+" && op2 == "*" {
		//	fmt.Println(a + b*c)
		return a + b*c
	} else if op1 == "+" && op2 == "/" {
		//	fmt.Println(a + b/c)
		return a + b/c
	} else if op1 == "-" && op2 == "+" {
		//	fmt.Println(a - b + c)
		return a - b + c
	} else if op1 == "-" && op2 == "-" {
		//	fmt.Println(a - b - c)
		return a - b - c
	} else if op1 == "-" && op2 == "*" {
		//	fmt.Println(a - b*c)
		return a - b*c
	} else if op1 == "-" && op2 == "/" {
		//	fmt.Println(a - b/c)
		return a - b/c
	} else if op1 == "*" && op2 == "+" {
		//	fmt.Println(a*b + c)
		return a*b + c
	} else if op1 == "*" && op2 == "-" {
		//	fmt.Println(a*b - c)
		return a*b - c
	} else if op1 == "*" && op2 == "*" {
		//	fmt.Println(a * b * c)
		return a * b * c
	} else if op1 == "*" && op2 == "/" {
		//	fmt.Println(a * b / c)
		return a * b / c
	} else if op1 == "/" && op2 == "+" {
		//	fmt.Println(a/b + c)
		return a/b + c
	} else if op1 == "/" && op2 == "-" {
		//	fmt.Println(a/b - c)
		return a/b - c
	} else if op1 == "/" && op2 == "*" {
		//	fmt.Println(a / b * c)
		return a / b * c
	} else if op1 == "/" && op2 == "/" {
		//	fmt.Println(a / b / c)
		return a / b / c
	} else {
		//	fmt.Println("not found")
		return 0
	}
}
