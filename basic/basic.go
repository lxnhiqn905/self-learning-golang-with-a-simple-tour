package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------")
	// 1. Package: main
	// 2. Import:
	// 3. Export name
	// 4. Functions
	test := add(3, 7)
	fmt.Println(test)

	// 6. Function multi result
	fmt.Println("---------------------")
	a, b := swap(1, 2)
	fmt.Println(a)
	fmt.Println(b)

	// 7. Function return by name
	fmt.Println("---------------------")
	a, b = cal(1, 1)
	fmt.Println(a)
	fmt.Println(b)

	// 8. Variables initializers
	fmt.Println("---------------------")
	var c, d int = 1, 2
	var java, python = "This is java", false
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(java)
	fmt.Println(python)

	// 9. Short variable declarations => Using := => Declaration k variable and init type
	fmt.Println("---------------------")
	k := 99
	fmt.Println(k)

	// 11. Basic types
	fmt.Println("---------------------")
	var Tobe bool = true
	fmt.Printf("This is %v\n", Tobe)

	// Var multi
	fmt.Println("---------------------")
	var (
		MaxInt    int    = 1000
		Text      string = "This is Text"
		Shortname        = "LX"
	)
	fmt.Println(MaxInt)
	fmt.Println(Text)
	fmt.Println(Shortname)

	// Zero values
	fmt.Println("---------------------")
	var iInt int       // Is 0
	var iString string // Is ""
	var iBool bool     // Is false
	fmt.Println(iInt)
	fmt.Println(iString)
	fmt.Println(iBool)

	// Constant
	fmt.Println("---------------------")
	const Pi = 3.14
	fmt.Println(Pi)

	// Numeric Constants
	fmt.Println("---------------------")
	const (
		Big     = 1 << 100
		Small   = Big >> 99
		NumberF = 20000
	)
	var BigF float64 = Big
	var SmallF int = Small
	fmt.Println(BigF)
	fmt.Println(SmallF)
	fmt.Println(NumberF)
	fmt.Println(needFloat(BigF))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// Numeric Constants
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func cal(a, b int) (ret1, ret2 int) {
	// 7. Function return by name
	ret1 = a - b
	ret2 = a + b
	return
}

func swap(a, b int) (int, int) {
	// 6. Function multi result
	return b, a
}

func add(a, b int) int {
	// 4. Functions
	return a + b
}
