package main

import (
	"fmt"
	"time"
)

func function() {
	for {
		fmt.Println("executes every 250ms")
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	go function()
	go func() {
		for {
			fmt.Println("executes every 500ms")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
}
