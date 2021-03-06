package main

import "fmt"

func main() {

	power2 := make(map[int]int)

	power2[2] = 4
	power2[3] = 9

	fmt.Println(power2)
	fmt.Println(power2[2])
	fmt.Println(power2[3])
	fmt.Println(power2[1]) // does not exist

	delete(power2, 1)
	delete(power2, 3)
	fmt.Println(power2)

	power2[4] = 16
	power2[3] = 9
	power2[5] = 25
	power2[5] = 25
	power2[5] = 55
	power2[10] = 25

	fmt.Println(power2)
	fmt.Println(len(power2))

}
