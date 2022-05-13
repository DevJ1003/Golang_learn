package main

import "fmt"

func main() {

	var radius float32 = 6
	fmt.Printf("The radius of a circle : %g\n", radius)

	var area float32 = ((radius * radius) * 22) / 7
	fmt.Printf("The area of a circle with radius %g : %g\n", radius, area)

}
