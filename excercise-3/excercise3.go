package main

import "fmt"

func main() {

	const MinRandonVal = 10
	const MaxrandonVal = 20

	var YourRandonVal int

	for {

		fmt.Printf("Enter values between %d and %d : ", MinRandonVal, MaxrandonVal)
		fmt.Scanf("%d", &YourRandonVal)

		if (YourRandonVal < MinRandonVal) || (YourRandonVal > MaxrandonVal) {

			fmt.Print("Incorrect value(s) entered, please try again...")
			// continue
		} else {
			break
		}

	}

}
