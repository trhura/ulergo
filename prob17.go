package main

import "fmt"

func main () {
	sum := 0
	for i := 1; i < 1000; i++ {
		sum += countLetters(getWrittenString(i))
	}
	sum += countLetters("one thousand")
	fmt.Println(sum)
}

func getWrittenString (number int) string {
	digits_to_str := map[int]string {
		0: "", // we don't need single zero
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}


	written_string := ""

	// write the string for third digit (hundred)
	third_digit := number / 100
	if third_digit > 0 {
		written_string += digits_to_str[third_digit] + " hundred"
	}

	last_two_digits := (number % 100)
	// add `and` if necessary
	if  number > 100 && last_two_digits > 0 {
		written_string +=  " and "
	}

	// write the string for the last two digits
	if (last_two_digits < 20 || (last_two_digits % 10) == 0) {
		written_string += digits_to_str[last_two_digits]
	} else {
		written_string += digits_to_str[(last_two_digits/10)*10] + "-" + digits_to_str[(last_two_digits%10)]
	}

	return written_string
}

func countLetters (str string) int {
	i := 0
	for _, c := range str {
		if c != rune('-') && c != rune(' ') { i++ }
	}
	return i
}
