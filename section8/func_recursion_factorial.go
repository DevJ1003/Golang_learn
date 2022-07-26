package main

import "fmt"

func main() {

	fmt.Println("=>", factorial(3))
	fmt.Println("=>", factorial(4))
	fmt.Println("=>", factorial(7))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	fmt.Print(n, " ") //for debugging purposes
	return n * factorial(n-1)
}
