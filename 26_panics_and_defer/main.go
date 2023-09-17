package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer panic("2")
	defer fmt.Println("3")
	panic("4")
}
