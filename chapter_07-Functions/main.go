package main

import (
	"fmt"
	"sort"
)

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Variadic Functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	fmt.Println(add(1, 2, 3))

	defer second()
	first()

	// q1) sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
	// sum(nn ...float64) int

	// q2) Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
	num := 12
	fmt.Println("q2)", num, "divided in half is even?", halfEvenOrOdd(num))

	// q3) Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	fmt.Println("q3) greatest num in list:", greatestNum(10, 9, 7))

	// q4) Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers
	nextOdd := makeOddGenerator()
	fmt.Println("q4)", nextOdd())
	fmt.Println("q4)", nextOdd())
	fmt.Println("q4)", nextOdd())
}

func halfEvenOrOdd(i int) bool {
	half := i / 2
	if half%2 == 0 {
		return true
	}
	return false
}

func greatestNum(ii ...int) int {
	find := func(i ...int) int {
		sort.Ints(i)
		return i[len(i)-1]
	}
	return find(ii...)
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
