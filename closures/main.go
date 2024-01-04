package main

import "fmt"

// closure definition: function that references variables outside its scope.

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println("pointer 1,", &i)
		return i
	}
}

func incrementor2() func() *int {
	i := 0
	return func() *int {
		i++
		fmt.Println("pointer 2,", &i)
		return &i
	}
}

func main() {

	inc := incrementor()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	// You can change the inner variable.
	inc2 := incrementor2()
	fmt.Println(*inc2())
	fmt.Println(*inc2())

	x := inc2()
	*x = 69
	fmt.Println(*inc2())
}
