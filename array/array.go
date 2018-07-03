package main

import (
	"fmt"
)

func main() {
	// Array
	fmt.Println("---------------------")
	var arrays [2]string // Definition array
	arrays[0] = "Nhi"
	arrays[1] = "LX"
	fmt.Println(arrays)    //All element
	fmt.Println(arrays[0]) // Each element
	fmt.Println(arrays[1]) // Each element

	arraysInt := [2]int{1, 20} // Define array and init value
	fmt.Println(arraysInt)     //All element

	var arraysStr = [2]string{"Test1", "Test2"} // Definition array init value
	fmt.Println(arraysStr)

	// Slice
	fmt.Println("---------------------")
	var sliceS []int
	arraysS := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceS = arraysS[3:6]
	sliceSView := arraysS[1:7]

	fmt.Println(sliceS)
	fmt.Println(sliceSView)

	sliceS[0] = 30
	sliceS[2] = 60

	fmt.Println(arraysS)
	fmt.Println(sliceSView)

	type iStruct struct {
		i int
		b bool
	}

	iArray := []iStruct{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(iArray)
	iSlice := iArray[2:4]
	fmt.Println(iSlice)
	fmt.Println(iSlice[0].i)
	fmt.Println(iSlice[0].b)
	iSlice[0].i = 15
	iSlice[1].i = 15
	fmt.Println(iArray)

	var nilSlice []int
	fmt.Println(nilSlice)

	if nilSlice == nil {
		fmt.Println("nil!")
	}
}
