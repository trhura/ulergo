package main

import "fmt"

func main() {
	const upperBound = 4000000
	var fib1 = 1
	var fib2 = 2
	var sum = 0

	for fib2 < upperBound {
		// add even fibonacci numbers
		if fib2 % 2 == 0 { sum += fib2 }
		fib1, fib2 = fib2, fib1
		fib2 = fib1 + fib2
	}

	fmt.Println(sum)
}
