package main

import "fmt"

func main() {

	const (
		No, False, Off = iota, iota, iota // all o
		Yes, True, On                     // all 1
	)

	fmt.Printf("No = %d False = %d Off = %d\n", No, False, Off)

	fmt.Printf("Yes = %d True = %d On = %d\n", Yes, True, On)

}
