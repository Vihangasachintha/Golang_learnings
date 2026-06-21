package main

import "fmt"

func main() {
	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)
	// fmt.Println(numbers[3])

	// var colors [3]string
	// colors[0] = "red"
	// colors[1] = "blue"
	// colors[2] = "green"

	// fmt.Println(colors)

	// colors[1] = "yellow"
	// fmt.Println(colors)

	// Two dimentional array

	var twoDArray = [3][3]int{{2, 4, 6}, {8, 10, 5}, {0, -7, 2}}
	fmt.Println(twoDArray)

	var newArray = [4][2]int {{1,3},{4,5},{7,8},{9,10}}
	fmt.Println(newArray)

	fmt.Println(newArray[0][0])
}
