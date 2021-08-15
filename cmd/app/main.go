package main

import (
	"os"
)

const minYear = 1900
const maxYear = 2999

type date struct {
	day int
	month int
	year int
}

// isLeapYear - checks if year is leap year
func isLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 && year % 400 != 0 {
			return false
		}
		return true
	}

	return false
}

func strNumToInt(num rune, multiplier int) int {
	return multiplier * int(byte(num)-'0')
}

func splitDate(date string) []int {
	var strSlice []string
	returnIntSlice := []int{}
	lastIdx := 0
	slashCount := 0
	for idx, element := range date {
		if byte(element) == 47 {
			if slashCount <= 1 {
				slashCount += 1
				strSlice = append(strSlice, date[lastIdx:idx])
				lastIdx = idx+1
			}
			if slashCount == 2 {
				strSlice = append(strSlice, date[lastIdx:])
			}
		}
	}
	
	numStore := 0
	for _, element := range strSlice{
		if len(element) > 1 {
			// Calculate days and months
			for dateIdx, i := range element{
				switch dateIdx {
				case 0:
					if byte(i) == 48 {
						continue
					} else {
						numStore = numStore + strNumToInt(i, 10)
					}
				case 1:
					numStore = numStore + strNumToInt(i, 1)
				}
			}
			// Calculate years
			if len(element) == 4 {
				numStore = 0
				for yearIdx, i := range element{
					switch yearIdx {
					// 1000x
					case 0:
						numStore = numStore + strNumToInt(i, 1000)
					// 100x
					case 1:
						numStore = numStore + strNumToInt(i, 100)
					// 10x
					case 2:
						numStore = numStore + strNumToInt(i, 10)
					// 1x
					case 3:
						numStore = numStore + strNumToInt(i, 1)
					}
				}
			}
		returnIntSlice = append(returnIntSlice, numStore)
		} else {
			numStore = numStore + strNumToInt(rune(element[0]), 1)
			returnIntSlice = append(returnIntSlice, numStore)
		}
		numStore = 0
	}

	return returnIntSlice
}

// initDate - will verify and init a new date structure via constructor pattern
func initDate(inputDate string) *date {
	// Split date - maybe verify here and pass error up
	intDateSlice := splitDate(inputDate)

	// Map valid days and months
	validDaysAndMonths := map[int][]int{
		1: {1, 31},
		2: {1, 28},
		3: {1, 31},
		4: {1, 30},
		5: {1, 31},
		6: {1, 30},
		7: {1, 31},
		8: {1, 31},
		9: {1, 30},
		10: {1, 31},
		11: {1, 30},
		12: {1, 31},
	}

	if isLeapYear(intDateSlice[2]) {
		validDaysAndMonths[2][1] = 29
	}
	if intDateSlice[2] <= minYear || intDateSlice[2] >= maxYear {
		panic("invalid year input")
	} 
	_, okMonth := validDaysAndMonths[intDateSlice[1]]
	if !okMonth {
		panic("invalid month input")
	}
	if intDateSlice[0] < validDaysAndMonths[intDateSlice[1]][0] || intDateSlice[0] > validDaysAndMonths[intDateSlice[1]][1] {
		panic("invalid day input")
	}

	d := new(date)
	d.day = intDateSlice[0]
	d.month = intDateSlice[1]
	d.year = intDateSlice[2]

	return d
}

func daysBetween(firstDate, secondDate date) int {
	// Map valid days and months
	validDaysAndMonths := map[int][]int{
		1: {1, 31},
		2: {1, 28},
		3: {1, 31},
		4: {1, 30},
		5: {1, 31},
		6: {1, 30},
		7: {1, 31},
		8: {1, 31},
		9: {1, 30},
		10: {1, 31},
		11: {1, 30},
		12: {1, 31},
	}

	dayCount := 0
	for {
		if firstDate.year == secondDate.year && firstDate.month == secondDate.month && firstDate.day == secondDate.day {
			dayCount--
			break
		}
		// Check leap year
		if isLeapYear(firstDate.year) {
			validDaysAndMonths[2][1] = 29
		}
		firstDate.day++
		// Increment month & reset day if last day is end of the month else increment day if day is over limit of month
		if firstDate.day > validDaysAndMonths[firstDate.month][1] {
			firstDate.month++
			firstDate.day = 1
		}
		// Increment year
		if firstDate.month > len(validDaysAndMonths) {
			firstDate.year++
			firstDate.month = 1
		}
		dayCount++
	}

	return dayCount
}

func main() {
	args := os.Args
	firstDate := initDate(args[1])
	secondDate := initDate(args[3])	
	println(daysBetween(*firstDate, *secondDate))
}
