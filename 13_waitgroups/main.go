package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	delay := rand.Float64() * 5
	time.Sleep(time.Duration(delay) * time.Second)

	if rand.Intn(1000) == 0 {
		fmt.Printf("Process %d slept for %.2f seconds\n", i, delay)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go process(wg, i)
	}

	wg.Wait()

	fmt.Println("All processes finished")
}
