package main

import "fmt"

func main() {

	var SeasonNo int
	fmt.Print("Enter the number for season : ")
	fmt.Scanf("%d", &SeasonNo)

	switch SeasonNo {

	case 1:

		fmt.Println("Spring - ", SeasonNo)
		// break
	case 2:
	case 3:
	case 4:
		fmt.Println("Summer - ", SeasonNo)
		fmt.Println("Rain - ", SeasonNo)
		fmt.Println("Cloudy - ", SeasonNo)

	default:

		fmt.Println("A new season - ", SeasonNo)

	}

}
