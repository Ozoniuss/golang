package aaabbb

// Note that in the importing package you should use the DIRECTORY NAME
// as the package name. However, the package is imported as the provided
// identifier here.

import "fmt"

func init() {
	fmt.Println("some_other_pkg was imported s.go")
}

func TryMeOut() {
	fmt.Println("try me out from some_other_pkg!")
}
