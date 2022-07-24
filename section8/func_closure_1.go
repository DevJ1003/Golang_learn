package main

import "fmt"

func main() {

	nextPos1 := getPossitiveInt()

	fmt.Println(nextPos1())
	fmt.Println(nextPos1())
	fmt.Println(nextPos1())

	fmt.Println()
	nextPos2 := getPossitiveInt()
	fmt.Println(nextPos2())

	fmt.Printf("'%T' '%T' \n", nextPos1, nextPos1())
	fmt.Printf("%x %x \n", &nextPos1, &nextPos2)

}

func getPossitiveInt() func() int {

	i := 0

	return func() int {
		i++
		return i
	}

}
