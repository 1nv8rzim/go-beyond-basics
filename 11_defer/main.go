package main

import "fmt"

func function() {
	defer fmt.Println("function() closed")

	fmt.Println("function() started")
}

func main() {
	defer fmt.Println("This runs after everything completes")

	fmt.Println("This is the first line")

	function()
}
