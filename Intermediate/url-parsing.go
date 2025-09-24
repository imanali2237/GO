// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// func main() {
// 	sampleUrl := "http://localhost:8080/path?name=john"
// 	parsedUrl, err := url.Parse(sampleUrl)
// 	if err != nil {
// 		fmt.Println("Error Occured While Parsing", err)
// 		return

// 	}
// 	fmt.Println("HOST", parsedUrl.Host)
// 	fmt.Println("PATH", parsedUrl.Path)
// 	// Getting value from Query
// 	queryParams := parsedUrl.Query().Get("name")
// 	fmt.Println(queryParams)
// 	// Building URL
// 	baseUrl := &url.URL{
// 		Scheme: "http",
// 		Host:   "example.com",
// 		Path:   "/Path",
// 	}
// 	query := baseUrl.Query()
// 	query.Set("key", "value")
// 	query.Set("name", "iman")
// 	baseUrl.RawQuery = query.Encode()
// 	fmt.Println("Hello", baseUrl)
// 	// Add Values to Query
// 	values := url.Values{}
// 	values.Add("age", "30")
// 	baseUrl.RawQuery = values.Encode()
// 	fmt.Println(baseUrl)

// }
