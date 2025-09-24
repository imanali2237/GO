// package main

// import "fmt"

// func main() {
// 	// result := calculateFactorial(4)
// 	// fmt.Println(result)

// 	// result := fibonacci(4)
// 	// fmt.Println(result)
// 	n := 10 // number of terms you want
// 	memo := make(map[int]int)

// 	fmt.Print("Fibonacci series: ")
// 	for i := 0; i < n; i++ {
// 		fmt.Print(fibonacci(i, memo), " ")
// 	}

// }
// func calculateFactorial(n int) int {
// 	// Base Case
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * calculateFactorial(n-1)

// }

// // Fibonacci Series

// // func fibonacci(number int) int {
// // 	// Base Case
// // 	if number <= 1 {
// // 		return number
// // 	}
// // 	return fibonacci(number-1) + fibonacci(number-2)

// // }

// // Fibonacci Series (Optimized)

// // fibonacci with recursion + memoization
// func fibonacci(n int, memo map[int]int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	if val, ok := memo[n]; ok {
// 		return val
// 	}
// 	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
// 	return memo[n]
// }
