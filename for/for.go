package main

import "fmt"

func main() {

	// iterator manually
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// auto iterator
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// infinite loop - break
	for {
		fmt.Println("loop")
		break
	}

	// loop - continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
