package main

import "fmt"

func main() {
	// Declare and initialize a map
	studentAges := map[string]int{
		"ajay": 23,
		"Dula": 23,
	}
	studentAges["Vihanga"] = 24
	studentAges["Tom"] = 12
	studentAges["Bob"] = 34

	studentAges["Bob"] = 3200

	fmt.Println(studentAges["Vihanga"])

	delete(studentAges, "Vihanga")
	fmt.Println(studentAges)

	fmt.Println(studentAges["Vihanga"])
}
