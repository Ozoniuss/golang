package main

// It was sufficient to add a replace directive in the go mod and I was able
// to import the local package.

import (
	"fmt"
	"icangiveitanyname/alphabetical"

	// "icangiveitanyname/executable" // CANNOT import executable packages, but can import if function is called main but package is called differently
	// "icangiveitanyname/internal" // CANNOT import internal package (containing internal in directory, not package name)
	other "icangiveitanyname/some_other_pkg"
	"icangiveitanyname/somepkg"
)

func main() {
	fmt.Println("importing")
	somepkg.TryMeOut()
	alphabetical.Alpha()
	other.TryMeOut()
	// fmt.Println(executable.EXECUTABLE)
	// fmt.Println(internal.IsCallable())
}
