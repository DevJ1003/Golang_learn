package main

import "fmt"

func main() {

	// BUBBLE SORT

	items := []int{18, 20, 14, 17, 13, 16, 11, 15, 12, 19}

	// order := 1 // 1 for ascending and -1 for descending

	end := len(items)

	for i := 0; i < end; i++ {
		for j := 0; j < (end - 1 - i); j++ {

			if items[j] > items[j+1] {

				// if ((items[j] - items[j + 1]) * order) > 1 {
				items[j], items[j+1] = items[j+1], items[j]

			}

		}

	}

	fmt.Printf("%v\n", items)

}
