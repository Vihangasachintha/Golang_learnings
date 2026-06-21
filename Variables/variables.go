package main

import "fmt"

func main() {

	var firstNumber int = 30
	var secondNumber int = 20
	var name string = "Vihanga"

	var userscore float32 = 25.5
	var userBalance float64 = -25.5425657

	fmt.Println(firstNumber + secondNumber)
	fmt.Println(firstNumber - secondNumber)
	fmt.Println(firstNumber * secondNumber)
	fmt.Println(firstNumber / secondNumber)

	fmt.Println(name)

	fmt.Println(userscore, userBalance)
	fmt.Printf("Type: %T\n", name) // Use to see data type of a var

	var isUserFound bool = false

	fmt.Println(isUserFound)

	var firstName, secondName = "Vihanga", "Sachintha"
	fmt.Println(firstName, secondName)

	newUserName := "Vihanga new"
	fmt.Println(newUserName)
	fmt.Printf("%T", newUserName)
}
