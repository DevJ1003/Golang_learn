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

		fmt.Println("Summer - ", SeasonNo)
		fallthrough
	case 3:

		fmt.Println("Rain - ", SeasonNo)
		fallthrough //it is optioanl
	case 4:

		fmt.Println("Cloudy - ", SeasonNo)

	default:

		fmt.Println("A new season - ", SeasonNo)

	}

}
