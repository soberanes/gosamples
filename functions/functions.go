package main

import "fmt"

// Functions are central in Go. We'll learn about functions with a
// few different examples.

// Plus is a function that takes two ints and returns their sum as an int
func plus(a int, b int) int {

	return a + b
}

// PlusPlus - when you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters
// up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

//  Main - Call a function just as you'd expect, with name(args)
func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

}
