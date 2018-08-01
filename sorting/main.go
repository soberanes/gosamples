package main

import (
	"fmt"
	"sort"
)

// Go's sort package implements sorting fo builtins and user-defined types.

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Integers:", ints)

	// We can also use sort to check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

}
