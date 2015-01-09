package main

import "fmt"

func main() {
	for i:= 20;; i += 2 {
		// check the number is evenly divisible by 1..20
		for num := 11; num <= 20; num++ {
			if i % num != 0 {
				break
			}

			if num == 20 {
				fmt.Println(i)
				return
			}
		}
	}
}
