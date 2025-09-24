// package basics

// import "fmt"

// func main() {
// 	deferFunction()

// }
// func deferFunction() {
// 	defer fmt.Println("Deferred Statement")         //this will execute at last
// 	defer fmt.Println("This will be first defered") //reason is defer follows LIFO principle who is defered First WIll be Executed at last
// 	fmt.Println("Normal Functions")                 //this will execute before defer
// }
