package main

import "fmt"

func main() {

	// 6th for loop programme - nested for loop statement

	for i := 1; i < 5; i++ {

		fmt.Printf("\ni = %d ** ", i)

		for j := 1; j < 3; j++ {

			fmt.Printf("%d ", i*j)

		}

	}

}
