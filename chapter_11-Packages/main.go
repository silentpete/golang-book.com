package main

import (
	"fmt"

	"./math"
	"./say"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	say.Durp("peter")
	say.Durp("new")
}
