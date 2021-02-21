package main

func main() {
	//calculateUsing3Operands()
	//calculateUsing3OperandsAndSameOp()
	//CalculateUsingSwitch(6, "+", 2)
	//calculateUsingIfelse()
	//CalculateUsingMap()
}
func CreateOperationMap() map[string]func(a, b int) int {
	return map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
		"%": func(a, b int) int { return a % b },
	}
}

func CreatePrecedenceMap() map[string]int {
	return map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
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
	p := CreatePrecedenceMap()
	if p[op1] < p[op2] {
		r := m[op2](b, c)
		return m[op1](a, r)
	}
	return m[op2](m[op1](a, b), c)
}

// func CalculateUsing4OperandsAndSameOp(a int, op1 string, b int, op2 string, c int, op3 string, d int) int {
// 	m := CreateOperationMap()
// 	r1 := m[op1](a, b)
// 	r2 := m[op2](r1, c)
// 	return m[op3](r2, d)
// }
//
// func CalculateUsing4Operands(a int, op1 string, b int, op2 string, c int, op3 string, d int) int {
// 	m := CreateOperationMap()
// 	if (op2 == "*" || op2 == "/") && (op3 == "*" || op3 == "/") && (op1 == "-" || op1 == "+") && (op1 == "+" || op2 == "+") {
// 		r1 := m[op3](c, d)
// 		r2 := m[op2](r1, a)
// 		return m[op1](b, r2)
// 	}
// 	// a op1 b op2 c op3 d
// 	// 6 + 4 * 4 * 2
// 	// 6 + 16 * 2
// 	// 6 + 32
// 	// 38
//
// 	//r1 := m[op1](a, b)
// 	//r2 := m[op2](r1, c)
// 	//return m[op3](r2, d)
// 	return m[op3](m[op2](m[op1](a, b), c), d)
// }
