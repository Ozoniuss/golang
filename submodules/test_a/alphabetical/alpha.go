package alphabetical

import "fmt"

// imported packages are loaded in alphabetical order

func init() {
	fmt.Println("alphabetical was imported")
}

func Alpha() string {
	return "alpha"
}
