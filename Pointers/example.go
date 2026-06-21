package main

import "fmt"

// Function to incremnent a value using a pointer
func incremnent(value *int) {
	*value++
}

// Function to set a value to zero using a pointer
func setToZero(value *int) {
	*value = 0
}

func main() {

	num := 10
	incremnent(&num)

	fmt.Println("After incrementing:", num)

	setToZero(&num)
	fmt.Println("After setting to zero:", num)

}
