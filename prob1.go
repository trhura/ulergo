package main

import "fmt"

func main() {
	var sum = 0
	const upperBound = 1000

	// multiplce of three
	for i := 0; i < upperBound; i += 3 {
		sum += i
	}

	// multiple of five
	for i := 0; i < upperBound; i += 5 {
		// don't count number already divided by  3
		if i % 3 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
