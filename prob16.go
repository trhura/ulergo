package main

import "math/big"
import "fmt"
import "strconv"

func main() {

	n := big.NewInt(1)
	for i := 0; i < 1000; i++ {
		n.Mul(n, big.NewInt(2))
	}

	sum := 0
	for _, c := range n.String() {
		i, _ := strconv.Atoi(string(c))
		sum += i
	}
	fmt.Printf ("sum = %d\n", sum)
}
