package main

import "fmt"

func main() {

	largest_chain_number := 0
	largest_chain_length := 0

	for i := 1; i < 1000000; i++ {
		n := i
		chain_length := 0

		for {
			chain_length++

			if n == 1 { break }
			if (n % 2) == 0 {
				n /= 2
			} else {
				n = (3*n) + 1
			}
		}

		if chain_length > largest_chain_length {
			largest_chain_length = chain_length
			largest_chain_number = i
		}
	}

	fmt.Println(largest_chain_number, largest_chain_length)
}
