package package1

import "fmt"

func init() {
	fmt.Println("init() in package1")
}

func Hello() {
	fmt.Println("Hello from package1")
}
