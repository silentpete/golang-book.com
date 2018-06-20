package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	var first string = "one way to create a variable"
	second := "second way to create a variable"
	fmt.Println(first)
	fmt.Println(second)

	x := 5
	x += 1
	fmt.Println(x)
}
