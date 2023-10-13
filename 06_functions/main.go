package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func createFile(name string) (*os.File, error) {
	file, err := os.Create(name)

	return file, err
}

func deleteFile(name string) error {
	return os.Remove(name)
}

func main() {
	solution := add(1, 2)
	fmt.Println(solution)

	solution = sum(1, 2, 3, 4, 5)
	fmt.Println(solution)

	_, err := createFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	deleteFile("test.txt")
}
