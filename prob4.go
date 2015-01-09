package main

import "fmt"
import "strconv"

func main() {
	l := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= i; j-- {
			if isPalindrome(strconv.Itoa(i * j)) {
				if  (i * j) > l {
					fmt.Println(i,"*", j, "=", i*j)
					l = i * j
				}
			}
		}
	}
}

func isPalindrome(s string) bool {
	ls := len(s)
	for i := 0; i <= (ls/2); i++ {
		if (s[i] != s[ls-i-1]) {
			return false;
		}
	}

	return true;
}
