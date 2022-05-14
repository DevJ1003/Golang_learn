package main

import "fmt"

func main() {

	var year int
	fmt.Print("\nEnter year : ")
	fmt.Scanf("%d", &year)

	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeapYear {

		fmt.Printf("\nYear %d is a leap year!\n", year)

	} else {

		fmt.Printf("\nYear %d is not a leap year!\n", year)

	}

}
