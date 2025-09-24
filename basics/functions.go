// package main

// import "fmt"

// func main() {

// 	// result := applyOperation(1, 2, add)
// 	// fmt.Println("Result of First", result)
// 	// multiply := createMultiplier(2)
// 	// fmt.Println(multiply(6))

// 	// q, r := divide(10, 3)
// 	// fmt.Println("Quotient", q, "Value", r)

// 	result := VariadicFunctions(1, 5, 6, 7)
// 	fmt.Println(result)

// }

// // Simple Functions
// // func add(a, b int) int {
// // 	return a + b
// // }

// // // Functions Takes another Function as Argument
// // func applyOperation(x int, y int, operation func(int, int) int) int {
// // 	return operation(x, y)
// // }

// // // Functions Returns Another Function

// // func createMultiplier(factor int) func(int) int {
// // 	return func(i int) int {
// // 		return factor * i
// // 	}
// // }

// // Multiple Return Type Functions
// // func divide(a, b int) (int, int) {
// // 	quotient := a / b
// // 	reminder := a % b
// // 	return quotient, reminder
// // }

// // Variadic Functions
// // Syntax func funcName(param1 type1,param2 ...type2)returnType{Body}

// func VariadicFunctions(a int, b ...int) int {
// 	sumOfSecondParameters := 0
// 	for _, v := range b {
// 		sumOfSecondParameters += v
// 	}
// 	return sumOfSecondParameters + a

// }
