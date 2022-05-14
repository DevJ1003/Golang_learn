package main

import (
	"fmt"
)

func main() {

	// totalMilliSeconds := time.Now().UnixNano() / 1000000
	// fmt.Println(totalMilliSeconds)

	totalMilliSeconds := 16523888
	fmt.Printf("totalMilliSeconds : %d\n", totalMilliSeconds)

	totalSeconds := totalMilliSeconds / 1000
	fmt.Printf("totalSeconds : %d\n", totalSeconds)

	currentSeconds := totalSeconds % 60
	fmt.Printf("currentSeconds : %d\n", currentSeconds)

	totalMinutes := totalSeconds / 60
	fmt.Printf("totalMinutes : %d\n", totalMinutes)

	currentMinutes := totalMinutes % 60
	fmt.Printf("currentMinutes : %d\n", currentMinutes)

	totalHours := totalMinutes / 60
	fmt.Printf("totalHours : %d\n", totalHours)

	currentHours := totalHours % 24
	fmt.Printf("currentHours : %d\n", currentHours)

}
