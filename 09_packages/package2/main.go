package package2

import "fmt"

func init() {
	fmt.Println("init() in package2")
}

func Hello() {
	fmt.Println("Hello from package2")
}
