package main

import "fmt"

func main() {

	var intNum int = 10
	var floatNum float64 = 3.5

	sum := intNum + int(floatNum)
	fmt.Println(sum)

	sum2 := float64(intNum) + floatNum
	fmt.Println(sum2)

	sum3 := 3
	fmt.Println(sum3)
}
