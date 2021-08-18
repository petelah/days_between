package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const minYear = 1900
const maxYear = 2999
const usage = "usage: [command] <date> to <date>. ie: ./main 2/6/1983 to 22/6/1983"

type date struct {
	day   int
	month int
	year  int
}

var validDaysAndMonths = map[int][]int{
	1:  {1, 31},
	2:  {1, 28},
	3:  {1, 31},
	4:  {1, 30},
	5:  {1, 31},
	6:  {1, 30},
	7:  {1, 31},
	8:  {1, 31},
	9:  {1, 30},
	10: {1, 31},
	11: {1, 30},
	12: {1, 31},
}

// isLeapYear - checks if year is leap year
func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}

	return false
}

// initDate - will verify and init a new date structure via constructor pattern
func initDate(inputDate string) (*date, error) {
	strSlice := strings.Split(inputDate, "/")
	// Split date - maybe verify here and pass error up
	var intDateSlice []int
	for _, element := range strSlice {
		intFromStr, strConvErr := strconv.Atoi(element)
		if strConvErr != nil {
			return nil, fmt.Errorf("invalid string input")
		}
		intDateSlice = append(intDateSlice, intFromStr)
	}

	if isLeapYear(intDateSlice[2]) {
		validDaysAndMonths[2][1] = 29
	}
	if intDateSlice[2] <= minYear || intDateSlice[2] >= maxYear {
		return nil, fmt.Errorf("invalid year input. Must be between 1900 & 2999")
	}
	_, okMonth := validDaysAndMonths[intDateSlice[1]]
	if !okMonth {
		return nil, fmt.Errorf("invalid month input. Input range 1-12")
	}
	if intDateSlice[0] < validDaysAndMonths[intDateSlice[1]][0] || intDateSlice[0] > validDaysAndMonths[intDateSlice[1]][1] {
		return nil, fmt.Errorf("invalid day input. Input range 1-31")
	}

	d := new(date)
	d.day = intDateSlice[0]
	d.month = intDateSlice[1]
	d.year = intDateSlice[2]

	return d, nil
}

// daysBetween calculates the days between the two dates given
func daysBetween(firstDate, secondDate date) (int, error) {
	// Check firstDate is smaller than secondDate, if so swap them
	if firstDate.year <= secondDate.year {
		if firstDate.month > secondDate.month {
			firstDate, secondDate = secondDate, firstDate
		}
		if firstDate.month <= secondDate.month {
			if firstDate.day > secondDate.day {
				firstDate, secondDate = secondDate, firstDate
			}
		}
	} else {
		// Base case where firstDate.year > secondDate.year
		firstDate, secondDate = secondDate, firstDate
	}

	dayCount := 0
	for {
		// Check leap year
		if isLeapYear(firstDate.year) {
			validDaysAndMonths[2][1] = 29
		} else {
			validDaysAndMonths[2][1] = 28
		}

		// Case for exiting the loops when dates match
		if firstDate == secondDate {
			dayCount--
			break
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

	return dayCount, nil
}

func main() {
	args := os.Args

	// Validate args with simple length test
	if len(args) < 4 {
		log.Fatal(fmt.Sprintf("invalid input error. %s", usage))
	}

	firstDate, firstDateErr := initDate(args[1])
	if firstDateErr != nil {
		log.Fatal(fmt.Sprintf("date input error: %s\n%s", firstDateErr, usage))
	}

	secondDate, secondDateErr := initDate(args[3])
	if secondDateErr != nil {
		log.Fatal(fmt.Sprintf("date input error: %s\n%s", secondDateErr, usage))
	}

	days, daysErr := daysBetween(*firstDate, *secondDate)
	if daysErr != nil {
		log.Fatal(fmt.Sprintf("error calculating difference in dates: %s\n%s", daysErr, usage))
	}

	fmt.Printf("%d days", days)
}
