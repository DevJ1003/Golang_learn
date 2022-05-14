package main

import (
	"fmt"
	"os"
)

func main() {

	var SeasonNo string

	// for { // } infinite loop
	for SeasonNo != "end" {

		fmt.Print("Enter your favourite season (type end to exit) : ")
		fmt.Scanf("%s", &SeasonNo)

		// if SeasonNo == "end" {
		// 	break
		// }

		switch SeasonNo {

		case "spring", "SPRING":

			fmt.Fprintf(os.Stdout, "You entered %s\n", SeasonNo)

		case "summer", "SUMMER":
			fmt.Println("You entered : ", SeasonNo)
			break
			fmt.Println("This line won't be reached", SeasonNo) //Note

		case "fall", "FALL":

			fmt.Println("You entered : ", SeasonNo)

		case "winter", "WINTER":

			fmt.Println("You entered : ", SeasonNo)

		default:

			fmt.Println("A new season - ", SeasonNo)

		}
	}
}
