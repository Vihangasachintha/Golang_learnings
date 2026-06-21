package main

import "fmt"

// func main() {

	var num int = 10
	fmt.Println(num)

	// Declare a pointer to an integer
	var ptr *int
	ptr = &num

	fmt.Println(ptr)

	// Pointer Dereference
	fmt.Println(*ptr)


	num = 100
	*ptr = 200

	fmt.Println(&ptr)
}
