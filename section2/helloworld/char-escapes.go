// single line comment

/*

multi line comment

*/

package main

import "fmt"

func main() {

	fmt.Println("Col-1\tCol-2\tCol-3")
	fmt.Println("--------------------------\n")
	fmt.Println("item11\titem21\titem31")
	fmt.Println("item41\t\titem81")

	// another way (complicated)
	fmt.Println("\n%s\t%s\t%s", "Col-1", "Col-2", "Col-3")
	fmt.Println("\n--------------------------")
	fmt.Println("\n%s\t%s\t%s", "item11", "item21", "item31")
	fmt.Println("\n%s\t%s\t%s", "item41", "", "item81")

}
