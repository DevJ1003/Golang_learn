package main

import (
	"fmt"
	"os"
)

func main() {

	nums := [...]int{10, 15, 30}

	var sum int

	nums[1] = 20

	for i := range nums {

		sum += nums[i]

	}

	fmt.Println(sum)
	fmt.Println(nums)
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	// fmt.Println(nums[3]) -- compiler error - out of bounds
	fmt.Println(len(nums))

	var j int
	fmt.Println()
	for j = 0; j < (len(nums)); j++ {
		fmt.Print(nums[j], " ")
	}

	fmt.Fprintf(os.Stdout, "\n%d\n\n", j)

	sums := []int{4, 5, 6}
	fmt.Println(sums)
	fmt.Printf("\n\n")
	fmt.Fprintln(os.Stdout, sums)

}
