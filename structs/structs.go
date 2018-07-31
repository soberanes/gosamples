package main

import "fmt"

// Go's structs are typed collections of fields. They're useful for
// grouping data together to form records

// Person struct type has name and age fields
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers
	// are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
