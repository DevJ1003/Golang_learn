package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "Hello"
	s2 := "象形字"

	// 'RuneCountInString' returns number of runes (characters in the case).
	// 'len' returns the number of bytes occupied by the string.

	fmt.Println(len(s1), utf8.RuneCountInString(s1))
	fmt.Println(len(s2), utf8.RuneCountInString(s2))

}
