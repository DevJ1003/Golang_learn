package main

import "fmt"

func main() {

	fmt.Println(2, 20)
	fmt.Println(3, 30)
	fmt.Println(4, 40)

	fmt.Println("For loop starts ============")

	// 1st for loop
	i := 2
	for i < 5 {

		fmt.Println(i, i*10)
		i++

	}

}
