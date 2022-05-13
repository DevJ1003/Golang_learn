package main

import "fmt"

func main() {

	s1 := "形声读"
	s2 := "象形字"

	var myRunes []rune

	for _, r := range s1 {

		myRunes = append(myRunes, r)

	}

	for _, r := range s2 {

		myRunes = append(myRunes, r)

	}

	fmt.Printf("%q\n", myRunes)
	fmt.Printf("%q\n", string(myRunes))

}
