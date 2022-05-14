package main

import "fmt"

func main() {

	// if statement
	isSunny := true

	if isSunny {

		fmt.Println("I will play cricket!")

	}

	// if-else statement
	isRain := false

	if isRain {

		fmt.Println("I will study!")

	} else {

		fmt.Println("I will play cricket!")

	}

	// if-else-if statement
	Season1 := false // rain
	Season2 := false // sunny

	if Season1 {

		fmt.Println("I will study!")

	} else if Season2 {

		fmt.Println("I will play cricket!")

	} else {

		fmt.Println("I will play carrom!")

	}

}
