package main

import (
	"fmt"
	"strings"
)

func main() {
	// Make slice
	slice1 := make([]int, 5)
	fmt.Println(slice1)

	slice2 := make([]int, 0, 5)
	fmt.Println(slice2)

	array1 := [5]int{1, 2, 3, 4, 5}
	slice2 = array1[1:3]
	fmt.Println(slice2)
	fmt.Println(slice2[0])
	slice2[0] = 1
	fmt.Println(slice2)

	// Slice of slice
	sliceofslice := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(sliceofslice)

	// The players take turns.
	sliceofslice[0][0] = "X"
	sliceofslice[2][2] = "O"
	sliceofslice[1][2] = "X"
	sliceofslice[1][0] = "O"
	sliceofslice[0][2] = "X"

	for i := 0; i < len(sliceofslice); i++ {
		fmt.Printf("%s\n", strings.Join(sliceofslice[i], " "))
	}

	// Append slice
	var silce3 []int
	fmt.Println(silce3)

	slice4 := append(silce3, 2)
	fmt.Println(slice4)
	slice4 = append(slice4, 3)
	fmt.Println(slice4)

	slice4 = append(slice4, 4, 5, 6, 6)
	fmt.Println(slice4)

	// Range
	var arrayNumber = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayNumber)

	for i, val := range arrayNumber {
		fmt.Printf("%d and %d----\n", i, val)
	}

	// Range with slice
	slice5 := make([]int, 10)
	fmt.Println(slice5)

	for index := range slice5 {
		slice5[index] = 1 << uint(index)
	}
	fmt.Println(slice5)

	// Map
	type vecTex struct {
		Lat, Long float32
	}
	var mapTex map[string]vecTex
	mapTex = make(map[string]vecTex)
	fmt.Println(mapTex)

	mapTex["1"] = vecTex{1.2, 3.4}
	fmt.Println(mapTex)

	type internet struct {
		author, link string
	}

	internetMap := map[string]internet{
		"facebook": internet{"facebook.com", "fb.com"},
		"google":   internet{"google.com", "gg.com"},
	}

	fmt.Println(internetMap)

	//Mutating Map
	mutaingmap := make(map[string]int)
	mutaingmap["10"] = 10
	mutaingmap["20"] = 20
	fmt.Println(mutaingmap)
	delete(mutaingmap, "20")
	fmt.Println(mutaingmap)

	v, ok := mutaingmap["10"]
	fmt.Println("The value:", v, "Present?", ok)

}
