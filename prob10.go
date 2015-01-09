package main

import "fmt"

func main () {
	sum := 0
	for i := 0; i < 2000000; i++ {
		if (isPrime(i)) {
			sum += i
		}
	}
	fmt.Println(sum)
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
