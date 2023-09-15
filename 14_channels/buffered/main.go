package main

func main() {
	buffered_channel := make(chan int, 1)

	buffered_channel <- 1
	buffered_channel <- 2 // deadlock

	println(<-buffered_channel)
}
