package main

import (
	"fmt"
	"time"
)

var delayTime = 750 * time.Millisecond // original = 100

func main() {

	executionTime := 5 * time.Second // original = 3
	start := time.Now()

	fmt.Printf("Program will end in about %v.\n", executionTime)
	fmt.Print("Operation in progress ... ")

	s := `\|/-`
	i := 0
	for {

		printSpinningSymbol(string(s[i]))

		if time.Since(start) > executionTime {

			fmt.Println("Done")
			fmt.Println("Elapsed Time : ", time.Since(start))
			break

		}

		i = (i + 1) % 4

	}

}

func printSpinningSymbol(symbol string) {

	fmt.Print(symbol)
	time.Sleep(delayTime)
	fmt.Print("\b")

}
