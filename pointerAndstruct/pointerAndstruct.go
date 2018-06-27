package main

import (
	"fmt"
)

func main() {
	// POINTER
	fmt.Println("--------------------------")
	i := 10
	p := &i         // Pointer p, save value i
	fmt.Println(p)  // Pointer p
	fmt.Println(*p) // Value in Pointer p - Is i
	*p = 20         // Re-assign new value for pointer p
	fmt.Println(*p)
	fmt.Println(i)      // Value of i also changed
	fmt.Println(*p / 4) // Div with value of Poiter

	// Struct
	fmt.Println("--------------------------")
	type vexter struct {
		x, Y int
	} // Define struct
	vexterI := vexter{}    // Init struct, value will 0, 0
	fmt.Println(vexterI.x) // Get x in struct
	fmt.Println(vexterI.Y) // Get Y in struct

	vexterII := vexter{1, 5} // Init and assign value for struct
	fmt.Println(vexterII.x)  // Get x in struct
	fmt.Println(vexterII.Y)  // Get Y in struct
	vexterII.x = 10          // Re-assign x = 10
	vexterII.Y = 50          // Re-assign y = 50
	fmt.Println(vexterII.x)  // Get x in struct
	fmt.Println(vexterII.Y)  // Get Y in struct

	// POINTER and Struct
	fmt.Println("--------------------------")
	type mapG struct {
		ido, kido int
	} // Define struct
	pMapG := &mapG{}    // Pointer pMapG store value mapG, init 0-0 value
	fmt.Println(pMapG)  // Show pointer
	fmt.Println(*pMapG) // Show pointer value

	ppMapG := &mapG{20, 40}     // Pointer pMapG store init value mapG
	fmt.Println(ppMapG)         // Show pointer
	fmt.Println(*ppMapG)        // Show pointer value
	fmt.Println((*ppMapG).ido)  // Show each value
	fmt.Println((*ppMapG).kido) // Show each value

	(*ppMapG).ido = 200         // Re-assign value
	(*ppMapG).kido = 400        // Re-assign value
	fmt.Println((*ppMapG).ido)  // Show each value
	fmt.Println((*ppMapG).kido) // Show each value

	// Struct assignment
	fmt.Println("--------------------------")
	mapGInitNoValue := mapG{} // Init no value 0-0
	fmt.Println(mapGInitNoValue)
	mapGInit1Value := mapG{ido: 99} // Init 1 value ido = 99, kido = 0 (default of int)
	fmt.Println(mapGInit1Value)
	mapGInit2Value := mapG{88, 77} // Init 2 value ido = 88, kido = 77
	fmt.Println(mapGInit2Value)
	mapGInitPointer := &mapG{66, 55} // Init pointer 2 value ido = 66, kido = 55
	fmt.Println(mapGInitPointer)

}
