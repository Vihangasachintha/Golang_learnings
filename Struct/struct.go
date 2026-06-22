package main

import "fmt"

// Define a struct type for person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	vihanga := Person{
		FirstName: "Vihanga",
		LastName:  "Sachintha",
		Age:       24,
	}
	// fmt.Printf("The first name of the person is : %v", vihanga.FirstName)

	fmt.Println(vihanga)

	vihanga.Age = 23

	fmt.Println(vihanga)
}
