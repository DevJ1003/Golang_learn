package main

import "fmt"

func main() {

	// 3rd for loop programme
	i := 2

	for {

		if i == 8 {

			break

		}

		if i == 4 || i == 6 {

			i++
			continue

		}

		fmt.Print(i, " ")
		i++

	}

}
