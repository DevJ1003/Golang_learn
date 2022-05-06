// single line comment

/*

multi line comment

*/

package main

import "fmt"

func main() {

	fmt.Println(len("Dog Day"))
	fmt.Println("Dog Day"[5])
	fmt.Println() //for making space between the lines

	s := "SuperHelpful"

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:9])
	fmt.Println(s[:])

	fmt.Println(s[5:] + s[:5])

	s += "Friend"
	fmt.Println(s)

}
