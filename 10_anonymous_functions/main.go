package main

func sort(numbers []int, function func(int) bool) {
	for _, number := range numbers {
		if function(number) {
			println(number)
		}
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	function := func(i int) bool {
		return i%2 == 0
	}

	sort(numbers, function)
}
