package main

import "fmt"

// Go supports pointers, allowing you to pass references to values
// and records within your program

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too
	fmt.Println("pointer:", &i)
}
