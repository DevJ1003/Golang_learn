package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const LowRange = 1
	const HighRange = 10
	const NumOfChances = 5

	rand.Seed(time.Now().UnixNano())

	RandomNo := LowRange + rand.Intn(HighRange-LowRange)

	// fmt.Println("** ", RandomNo)

	fmt.Printf("You have %d chances to guess "+"a system-generated random number.\n", NumOfChances)

	var YourNum int
	message := "You were not lucky this time. Try again!"

	for i := 0; i < NumOfChances; i++ {

		fmt.Printf("Enter a number between %d and %d (try #%d) : ", LowRange, HighRange, i+1)
		fmt.Scanf("%d", &YourNum)

		if YourNum == RandomNo {

			message = "You guessed the number!"
			break

		} else {

			if YourNum > RandomNo {

				fmt.Println("Pick a smaller number!")

			} else {

				fmt.Println("Pick a bigger number!")

			}

		}

	}

	fmt.Println("\n=================================================")
	fmt.Printf("======Your Number=%d Random Number=%d \n====== %s", YourNum, RandomNo, message)
	fmt.Println("\n=================================================")

}
