package main

import "fmt"

func main() {

	var f []float32

	f = append(f, 1.2)
	f = append(f, 2.4, 4.8, 8.16)

	fmt.Println(f)

	f = append(f, f...)
	fmt.Println(f)

}
