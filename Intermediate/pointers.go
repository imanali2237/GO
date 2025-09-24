// package main

// import "fmt"

// func main() {

// 	// Delcaration
// 	var ptr *int
// 	// Initialization
// 	var a int = 10
// 	ptr = &a //ptr  now points to  a's memory Address //REFERENCING

// 	fmt.Println("Memory Address of a", &a)
// 	fmt.Println("Memory Address to which ptr points to", ptr)
// 	// Both Addresses Same

// 	// Dereferencing a pointer means get the value of that memory address, means that is stored at that memory address
// 	// fmt.Println(*ptr)
// 	modifyValue(ptr)
// 	fmt.Println(a)

// }

// func modifyValue(ptr *int) {
// 	*ptr++
// }
