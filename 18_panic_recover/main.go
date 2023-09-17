package main

import "fmt"

func panic_and_recover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()
	panic("PANIC")

	fmt.Println("This does not run")
}

func main() {
	panic_and_recover()
	fmt.Println("After panic")
}
