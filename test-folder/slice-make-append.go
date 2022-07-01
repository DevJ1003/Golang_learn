package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	// nums := []int{11, 12, 13, 14, 15, 16}
	fmt.Println(nums)
	fmt.Println("\n\n")

	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)
	fmt.Println("\n\n")

	s2 := make([]int, 3)
	copy(s2, nums)
	fmt.Println(s2)
	fmt.Println("\n\n")

	s3 := append(s2, 8, 9)
	fmt.Println(s3)
	fmt.Println("\n\n")

	s4 := s2
	fmt.Println(s4)
	fmt.Println("\n\n")

	s5 := make([]int, 2, 3) // length and capacity printed
	fmt.Println(s5)
	fmt.Println("\n\n")

}
