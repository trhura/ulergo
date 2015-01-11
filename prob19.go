package main

import "fmt"

func main () {
	startingDay := 1
	totalSundays := 0
	for i := 1901; i <= 2000; i++ {
		totalSundays += countSundaysInFirstOfMonth(i, startingDay)
		fmt.Println(i, startingDay, countSundaysInFirstOfMonth(i, startingDay))
		startingDay = findStartingWeekdayforNextYear(i, startingDay)
	}

	fmt.Println(totalSundays)
}

func countSundaysInFirstOfMonth (year int, startingWeekday int) int {
	startWithSunday := 0
	startingDay := startingWeekday
	for i := 1; i <= 12; i++ {
		if startingDay == 6 { startWithSunday++ }
		startingDay = findStartingWeekDayforNextMonth(year, i, startingDay)
	}
	return startWithSunday
}

func findStartingWeekdayforNextYear(thisYear int, startingWeekday int) int {
	return ((findTotalDaysInYear(thisYear) % 7) + startingWeekday) % 7
}

func findStartingWeekDayforNextMonth(thisYear int, thisMonth int, startingWeekday int) int {
	return ((findTotalDaysInMonth(thisYear, thisMonth) % 7) + startingWeekday) % 7
}

func findTotalDaysInMonth (year int, month int) int {
	daysByMonth := map[int]int {
		1: 31,
		2: 28,
		3: 31,
		4: 30,
		5: 31,
		6: 30,
		7: 31,
		8: 31,
		9: 30,
		10: 31,
		11: 30,
		12: 31,
	}

	if month == 2 && isLeapYear(year) {
		return 29
	}

	return daysByMonth[month]
}

func findTotalDaysInYear (year int) int {
	if isLeapYear(year) { return 366 }
	return 365
}

func isLeapYear (year int) bool {
	if year % 4  == 0 {
		if year % 100 == 0 && year % 400 != 0 {
			return false
		}
		return true
	}
	return false
}
