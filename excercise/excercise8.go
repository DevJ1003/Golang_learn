package main

import "fmt"

func main() {

	const (
		c1 = iota
		c2 = 5
		c3 = iota
		c4 = 8
		c5 = iota
	)

	fmt.Printf("c1 = %d\nc2 = %d\nc3 = %d\nc4 = %d\nc5 = %d\n", c1, c2, c3, c4, c5)

}
