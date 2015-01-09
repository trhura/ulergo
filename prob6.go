package main

import "fmt"

func main() {

	sum_of_square := 0
	square_of_sum := 0

	for i := 1; i <= 100; i++ {
		square_of_sum += i
		sum_of_square += i * i
	}
	square_of_sum *= square_of_sum

	fmt.Println(square_of_sum - sum_of_square)
}
