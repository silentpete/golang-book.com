package main

import (
	"fmt"
	"sort"
)

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	y := []float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(y)
	fmt.Println("avg:", (total / float64(len(y))))

	var ary [5]int
	slice := make([]int, 5)

	fmt.Printf("%T\n", ary)
	fmt.Println(len(ary))
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println(elements["Li"])

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// How do you access the 4th element of an array or slice?
	q1Slice := []int{0, 1, 2, 3, 4}
	fmt.Println("q1) fourth element:", q1Slice[4])

	// What is the length of a slice created using: make([]int, 3, 9)?
	q2Slice := make([]int, 3, 9)
	fmt.Println("q2) length:", len(q2Slice))
	fmt.Println("q2) capacity:", cap(q2Slice))

	// given x := [6]string{"a","b","c","d","e","f"}: what would x[2:5] give you?
	q3Array := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("q3) i2 through i5-1:", q3Array[2:5])

	// Write a program that finds the smallest number in this list:
	q4Slice := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sort.Ints(q4Slice)
	fmt.Println("q4) smallest number in slice:", q4Slice[0])
}
