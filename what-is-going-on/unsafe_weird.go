package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uint8 = 123
	p := unsafe.Pointer(&i)
	i2 := *(*uint16)(p)
	// Adding the print changes the result.
	// fmt.Println(123 * 256)
	fmt.Println(i2)
}
