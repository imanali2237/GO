// package main

// import "fmt"

// func main() {
// 	// var numbers []int
// 	// var numbers1 = []int{1, 2, 3}
// 	// numbers2 := []int{2, 3, 4}
// 	// slice := make([]int, 5)
// 	a := [5]int{1, 2, 4, 5, 6}
// 	makeSliceFromArray := a[1:4]
// 	// fmt.Println(makeSliceFromArray)

// 	copySlice := make([]int, len(makeSliceFromArray))
// 	copy(copySlice, makeSliceFromArray)
// 	// fmt.Println(copySlice)
// 	// var nilSlice [] int
// 	// Iterating Over Slices
// 	for i, v := range copySlice {
// 		fmt.Println(i, v)
// 	}

// 	// Multi-Dimmensional-Slice

// 	twoDimmensionalSlice := make([][]int, 3)
// 	for i := 0; i < 3; i++ {
// 		innerlength := i + 1
// 		twoDimmensionalSlice[i] = make([]int, innerlength)
// 		for j := 0; j < innerlength; j++ {
// 			twoDimmensionalSlice[i][j] = i + j

// 		}

// 	}
// }
