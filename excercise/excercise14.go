package main

import "fmt"

var a = 1

func main() {

	{

		fmt.Println("Loc 1 - ", a)
		a := 4
		fmt.Println("Loc 2 - ", a)

	}

	a := 2
	fmt.Println("Loc 3 - ", a)
	a += 10

	{

		a += 50
		fmt.Println("Loc 4 - ", a)

	}

	fmt.Println("Loc 5 - ", a)
	changeA()

}

func changeA() {

	fmt.Println("Loc 6 - ", a)

	{

		a := 100
		fmt.Println("Loc 7 - ", a)

	}

	fmt.Println("Loc 8 - ", a)

}
