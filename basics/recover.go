// package main

// import "fmt"

// func main() {
// 	procss()
// 	fmt.Println("Returned From Process")

// }
// func procss() {
// 	defer func() {
// 		// if r := recover(); r != nil {
// 		// 	fmt.Println("Recovered", r)

// 		// }
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("Recovered", r)

// 		}

// 	}()
// 	fmt.Println("Start Process")
// 	panic("Something went Wrong")
// 	fmt.Println("End Process") //Unreachable code
// }
