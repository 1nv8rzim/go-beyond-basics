package main

import (
	"math/rand"
	"time"
)

func goroutine(c chan float64) {
	i := rand.Float64() * 5
	time.Sleep(time.Duration(i) * time.Second)
	c <- i
}

func main() {
	c := make(chan float64)
	for i := 0; i < 5; i++ {
		go goroutine(c)
	}

	for i := 0; i < 5; i++ {
		println(<-c)
	}
}
