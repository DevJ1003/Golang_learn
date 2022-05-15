package main

import "fmt"

func main() {

	const NumOfRows = 2
	const NumOfCols = 3

	fmt.Printf("\n%6s", " ")
	for j := 1; j < NumOfCols; j++ {

		fmt.Printf("%3d", j)

	}

	fmt.Print("\n= = = =")
	for j := 1; j <= NumOfCols; j++ {

		fmt.Printf("%2s ", "=")

	}

	for i := 1; i <= NumOfRows; i++ {

		fmt.Printf("\n%3d | ", i)
		for j := 1; j <= NumOfCols; j++ {

			fmt.Printf("%3d", i*j)

		}

	}

	fmt.Println()

}
