package main

import "fmt"

func main() {

	defer func() {
		errMsg := recover()
		fmt.Println("wewewewe")
		fmt.Println(errMsg)
	}()

	// Only use one of the following at the same time.

	var x map[string]int
	x["key"] = 10
	fmt.Println(x)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	panic("BOO...!!")
	fmt.Println("This line won't be printed!")

}
