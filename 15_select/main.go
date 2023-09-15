package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan3 := make(chan int)
	chan2 := make(chan int)

	for _, c := range []chan int{chan1, chan3, chan2} {
		go func(channel chan int) {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			channel <- 1
		}(c)
	}

	select {
	case <-chan1:
		fmt.Println("chan1")
	case <-chan2:
		fmt.Println("chan2")
	case <-chan3:
		fmt.Println("chan3")
	}
}
