package main

import "fmt"

// Unless you import test_a directly this is not loaded
func init() {
	fmt.Println("do you think this is loaded?")
}

func Test_a() {
	fmt.Println("testing a")
}

func main() {
	fmt.Println("hsaha")
}
