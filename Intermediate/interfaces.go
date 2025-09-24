// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geometry interface {
// 	area() float32
// 	perim() float32
// }
// type circle struct {
// 	radius float32
// }
// type rectangle struct {
// 	height, width float32
// }

// func (r rectangle) area() float32 {
// 	return r.height + r.width
// }
// func (r rectangle) perim() float32 {
// 	return 2 * (r.height + r.width)
// }
// func (r circle) area() float32 {
// 	return math.Pi + r.radius + r.radius
// }
// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// func main() {
// 	c := circle{radius: 2}
// 	r := rectangle{width: 2, height: 2}
// 	measure(r) //works fine as rectangle implements all function of geometry type so it is indeed of geometry type
// 	measure(c) //Here measure function accept only geometry type arguments but since circle implements only one function of geometry interface, it is not considered to be geometry type

// }
