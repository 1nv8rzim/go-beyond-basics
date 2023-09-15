package main

import (
	"fmt"
	"time"
)

func main() {
	exit := true
	for exit {
		fmt.Println("executes once")
		exit = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("executes twice")
	}

	for {
		fmt.Println("executes forever")
		time.Sleep(1 * time.Second)
	}

}
