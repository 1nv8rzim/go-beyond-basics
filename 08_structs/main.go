package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type People []Person

func (p Person) String() string {
	return fmt.Sprintf("Person[name=\"%s\", age=%d]", p.Name, p.Age)
}

func (p People) String() string {
	var s string = "People[\n"
	for _, person := range p {
		s += fmt.Sprintf("\t%s,\n", person)
	}
	s += "]"
	return s
}

func main() {
	people := People{
		Person{Name: "John", Age: 30},
		Person{Name: "Jane", Age: 25},
		Person{Name: "Bob", Age: 50},
	}

	fmt.Println(people)
}
