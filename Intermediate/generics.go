// // Make a Stack Data Structure Using Generics
// package main

// import "fmt"

// type Stack[T any] struct {
// 	elements []T
// }

// func (s *Stack[T]) push(element T) {
// 	s.elements = append(s.elements, element)

// }
// func (s *Stack[T]) pop() (T, bool) {
// 	if len(s.elements) == 0 {
// 		var zero T
// 		return zero, false

// 	}
// 	element := s.elements[len(s.elements)-1]
// 	s.elements = s.elements[:len(s.elements)-1]
// 	return element, true

// }
// func (s Stack[T]) printAll() {
// 	for _, v := range s.elements {
// 		fmt.Println(v)
// 	}
// }
// func (s *Stack[T]) isEmpty() bool {
// 	return len(s.elements) == 0

// }

// func main() {
// 	stack := Stack[int]{}
// 	stack.push(2)
// 	stack.push(3)
// 	stack.printAll()
// 	stack.pop()
// 	fmt.Println(stack.isEmpty())
// 	stack.printAll()

// }
