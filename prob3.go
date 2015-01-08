package main

import "fmt"

func main() {
	factors := make([]int, 0)
	number := 600851475143

	for {
		prime := findDivisiblePrime(number)
		factors = append(factors, prime)

		if number == prime { break }
		number = number / prime
	}

	fmt.Println(factors)
}

func findDivisiblePrime (a int) int {
	// return the smallest prime that can divide a
	prime := 2
	for prime < a {
		if (a % prime == 0) { return prime }
		prime = nextPrime(prime)
	}
	return a
}

func nextPrime(a int) int {
	// return prime number larger than a
	for  {
		a++;
		if (isPrime(a)) { return a }
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
