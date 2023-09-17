package main

import (
	"embed"
	"fmt"
)

//go:embed embed.txt
var embededString string

//go:embed embed
var embededFS embed.FS

func main() {
	fmt.Println(embededString + "\n")

	files, _ := embededFS.ReadDir("embed")
	for _, file := range files {
		fmt.Println(file.Name())
		buffer, _ := embededFS.ReadFile("embed/" + file.Name())
		fmt.Println(string(buffer))
	}
}
