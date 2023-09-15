package main

import "fmt"

func main() {
	Slice := []string{"a", "b", "c"}
	for index, value := range Slice {
		fmt.Printf("Slice[%d] = \"%s\"\n", index, value)
	}
	fmt.Println()

	Map := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range Map {
		fmt.Printf("Map[%s] = %d\n", key, value)
	}
	fmt.Println()

	Chan := make(chan int)
	go func() {
		Chan <- 1
		Chan <- 2
		Chan <- 3
		close(Chan)
	}()
	for value := range Chan {
		fmt.Printf("Chan = %d\n", value)
	}
}
