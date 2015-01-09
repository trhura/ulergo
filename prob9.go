package main

import "fmt"

func main() {
	u := 1000
	for c := u; c > 0; c-- {
		for a := 1; a < c && a < u - c - a ; a++ {
			b:= (u - a - c)
			if b > c { continue }

			if ((a*a) + (b*b)) == (c*c) {
			  	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
				fmt.Printf("a*b*c=%d\n", a*b*c)
			}
		}
	}
}
