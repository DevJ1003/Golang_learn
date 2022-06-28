package main

import "fmt"

func main() {

	s := "ABCDEFGHIJKL"

	slc := s[1:5]

	fmt.Printf("s=%v type(s)=%T slc=%v type(slc)=%T\n", s, s, slc, slc)

}
