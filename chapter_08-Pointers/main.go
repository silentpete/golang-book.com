package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func word(yPtr *string) {
	*yPtr = "peter"
}

func square(x *float64) {
	*x = *x * *x
}

func swap(aPtr, bPtr *int) {
	c := *aPtr
	d := *bPtr
	*aPtr = d
	*bPtr = c
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	name := "dude"
	word(&name)
	fmt.Println(name)

	z := 1.5
	square(&z)
	fmt.Println(z)

	// swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1)
	a := 1
	b := 2
	fmt.Printf("a=%v\nb=%v\n", a, b)
	swap(&a, &b)
	fmt.Printf("a=%v\nb=%v\n", a, b)
}
