package main

import (
	"fmt"
)

func main() {

	var fahrenheit float32 = 80

	fmt.Printf("Temperature in fahrenheit : %g\n", fahrenheit)

	var celsius float32 = ((fahrenheit - 32) * 5) / 9

	fmt.Printf("Temperature in celsius : %g\n", celsius)

}
