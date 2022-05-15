package main

import "fmt"

func main() {

	var year int

	fmt.Print("Enter the year to find chinese zodiac sign : ")
	fmt.Scanf("%d", &year)

	// List of animals : Rooster, Dog, Pig, Rat, Ox, Tiger, Rabbit, Dragon, Snake, Horse, Goat, Monkey
	switch year % 12 {

	case 0:
		fmt.Printf("Year %d is Rooster's year!", year)

	case 1:
		fmt.Printf("Year %d is Dog's year!", year)

	case 2:
		fmt.Printf("Year %d is Pig's year!", year)

	case 3:
		fmt.Printf("Year %d is Rat's year!", year)

	case 4:
		fmt.Printf("Year %d is Ox's year!", year)

	case 5:
		fmt.Printf("Year %d is Tiger's year!", year)

	case 6:
		fmt.Printf("Year %d is Rabbit's year!", year)

	case 7:
		fmt.Printf("Year %d is Dragon's year!", year)

	case 8:
		fmt.Printf("Year %d is Snake's year!", year)

	case 9:
		fmt.Printf("Year %d is Horse's year!", year)

	case 10:
		fmt.Printf("Year %d is Goat's year!", year)

	case 11:
		fmt.Printf("Year %d is Monkey's year!", year)

	}

}
