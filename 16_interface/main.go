package main

import "fmt"

type Printable interface {
	String() string
}

type Person struct{ Name string }

func (p Person) String() string { return p.Name }

type Book struct{ Title string }

func Print(p Printable) {
	fmt.Println(p)
}

func main() {
	Print(Person{"John"})
	Print(Book{"Harry Potter"})
}
