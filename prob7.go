package main

import "fmt"

func main() {
	nth := 10001
	n := 0

	for i := 1;; i++ {
		if isPrime(i) {
			n++

			if n == nth {
				fmt.Println(i)
				return
			}
		}
	}
}

func isPrime(a int) bool {
	// check whether a is prime
	if (a <= 1) { return false }
	if (a == 2) { return true }
	if (a % 2 == 0) { return false }

	for i := 3; i < a; i += 2 {
		if (a % i == 0) { return false }
		if (i*i > a) { break }
	}

	return true
}
