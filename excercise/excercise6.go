package main

import "fmt"

func main() {

	fmt.Println("Sunday=0 \nMonday=1 \nTuesday=2 \nWednesday=3 \nThursday=4 \nFriday=5 \nSaturday=6 \n")

	today := 0     // value for today can be any i.e 0-6
	duration := 85 // the day we want to find
	futureDay := (today + duration) % 7

	fmt.Printf("today = %d, duration = %d, futureDay = %d\n", today, duration, futureDay)

}
