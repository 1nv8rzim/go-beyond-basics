package main

type Count[T comparable] struct {
	Key   T
	Value int
}

type Counter[T comparable] []Count[T]

func (c Counter[T]) Get(key T) int {
	for _, count := range c {
		if count.Key == key {
			return count.Value
		}
	}
	return -1
}

func (c Counter[T]) Increment(key T) {
	for i, count := range c {
		if count.Key == key {
			c[i].Value += 1
			return
		}
	}
	c = append(c, Count[T]{key, 1})
}

func main() {
	counter := Counter[string]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	counter.Increment("b")

	println(counter.Get("b"))
}
