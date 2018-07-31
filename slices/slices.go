package main

import "fmt"

func main() {

	// Slices are a key data type in Go, giving a more powerful
	// interface to sequences than arrays

	// Unlike arrays, slices are typed only by the elements they
	// contain (not the number of elements)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// Slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy'd (the content only).
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high]
	l := s[2:5]
	fmt.Println("sl1", l)

	// the slices up to (but excluding)
	l = s[:5]
	fmt.Println("sl2", l)

	// and this slice up from (and including)
	l = s[2:]
	fmt.Println("sl3", l)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
