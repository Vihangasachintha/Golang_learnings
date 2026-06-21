package main

import "fmt"

const pi float32 = 3.14

func main() {

	var radius float32 = 5.0

	var area float32 = 0.0

	area = pi * radius * radius

	fmt.Printf("Area is: %.2f", area)

}
