package main

import "fmt"

func main() {

	var x interface{}

	// try other values for x : 10, 10.25, true, "hello", mult, no value (nil), uint(10)
	x = mult

	switch i := x.(type) {

	case int:
		fmt.Printf("%T %v", i, i)

	case float64:
		fmt.Printf("%T %v", i, i)

	case bool, string:
		fmt.Println("Type is bool or string!")

	case func(int) float64:
		fmt.Printf("%T", i)

	case nil:
		fmt.Println("x is nil!")

	default:
		fmt.Println("Type unknown...!!")

	}

}

func mult(i int) float64 {

	return float64(i) * 1.5

}
