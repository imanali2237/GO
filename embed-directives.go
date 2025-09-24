// // Embed Directives helps us to add files that are not with the extension go but needs to be in our executable
// package main

// import (
// 	"embed"
// 	_ "embed"
// 	"fmt"
// 	"io/fs"
// )

// // go:embed example.txt
// var content string //Storing content of our file (example.txt)

// // Embedding Directories
// //
// //go:embed basics
// var content2 embed.FS

// func main() {
// 	fmt.Println("Embed Content", content)
// 	file, err := content2.ReadFile("basics/noice.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(file))
// 	err = fs.WalkDir(content2, "basics", func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		fmt.Println(path)
// 		return nil

// 	})
// 	if err != nil {
// 		fmt.Println(err)

// 	}

// }
 