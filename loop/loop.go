package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// FOR classic
	fmt.Println("---------------------")
	for index := 0; index < 2; index++ {
		fmt.Printf("This is %v\n", index)
	}

	// FOR, not need init value
	fmt.Println("---------------------")
	sum := 1
	for sum < 8 {
		fmt.Printf("This is %v\n", sum)
		sum += sum
	}

	// FOR range
	nums := []int{2, 3, 4}
	fmt.Println("---------------------")
	for num := range nums {
		fmt.Printf("This is num %v\n", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("This is key %v\n", key)
		fmt.Printf("This is value %v\n", value)
	}

	// IF - ELSE
	fmt.Println("---------------------")
	if true {
		fmt.Printf("This is true %v\n", true)
	}

	var testV int = 1
	for testV < 5 {
		if testV != 1 {
			fmt.Printf("This is not 1 %v\n", testV)
		} else {
			fmt.Printf("This is 1 %v\n", testV)
		}
		testV++
	}

	// SWITCH
	fmt.Println("---------------------")
	testW := 1
	switch testW {
	case 1:
		fmt.Printf("This is 1 %v\n", testW)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// Switch evaluation order
	fmt.Println("---------------------")
	fmt.Println("Today is Saturday ???")
	switch time.Saturday {
	case time.Now().Weekday() + 0:
		fmt.Println("Today")
	case time.Now().Weekday() + 1:
		fmt.Println("Today + 1")
	default:
		fmt.Println("Today + multi!!!")
	}

	// Switch with no condition
	fmt.Println("---------------------")
	switch {
	case time.Now().Weekday().String() == "Thursday":
		fmt.Println("Today is not", time.Thursday)
	default:
		fmt.Println("Not know !!!")
	}

	// DEFER
	fmt.Println("---------------------")
	defer fmt.Println("Setting is 0 - Will run after func returned") // Regist defer first will run at second
	fmt.Println("Setting is 1 - Will run before defer")
	fmt.Println("Setting is 2 - Will run before defer - 1")
	fmt.Println("Setting is 3 - Will run before defer - 2")

	for deferI := 0; deferI < 10; deferI++ {
		defer fmt.Printf("This is value %v\n", deferI) // Regist defer second will run at first
	}
	fmt.Println("Finished for statement")
}
