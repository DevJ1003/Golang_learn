package main

import "fmt"

func main() {

	s := "Hi"
	s = "Hello"

	fmt.Println(s == "Hello")
	fmt.Println(s != "Hello")

	var (
		name1 = "Zoey"
		name2 = "Deo"
		name3 = "Joshi"
	)

	fmt.Println(name1, name2, name3)

}
