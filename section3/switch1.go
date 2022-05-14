package main

import "fmt"

func main() {

	SeasonNo := 5

	switch SeasonNo {

	case 1:

		fmt.Println("Spring - ", SeasonNo)
		// break
	case 2:

		fmt.Println("Summer - ", SeasonNo)

	case 3:

		fmt.Println("Rain - ", SeasonNo)

	case 4:

		fmt.Println("Cloudy - ", SeasonNo)

	default:

		fmt.Println("A new season - ", SeasonNo)

	}

}
