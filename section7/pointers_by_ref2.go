package main

import "fmt"

func main() {

	a := 1
	fmt.Printf("(1) a=%d &a=%x \n\n", a, &a)

	f3(&a)
	fmt.Printf("(2) a=%d \n\n", a) // SAME

	fmt.Printf("(3) a=%d \n", f3(&a)) // SAME

}

func f3(y *int) int {

	*y++
	fmt.Printf("(f) *y=%d y=%x \n", *y, y)
	return *y

}
