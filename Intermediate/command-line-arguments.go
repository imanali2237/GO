package main

import (
	"flag"
	"fmt"
)

func main() {
	// fmt.Println("Command", os.Args[0])
	// for i, args := range os.Args {
	// 	fmt.Println(i, args)
	// }
	// Using Flags
	var name string
	var gender string
	flag.StringVar(&name, "userName", "NoName", "User name")
	flag.StringVar(&gender, "userGender", "NO", "User Gender")
	flag.Parse()
	fmt.Println(name)
	fmt.Println(gender)
}
