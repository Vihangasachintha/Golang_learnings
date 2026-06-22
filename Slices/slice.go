package main

import "fmt"

func main() {
	//declare & initialize a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	// // Append elements to the slice
	// numbers = append(numbers, 6, 7, 8)

	numbers = append(numbers[:4], numbers[5:]...)

	fmt.Println(numbers)
}
