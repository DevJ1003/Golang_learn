package main

import "fmt"

// TO-DO : write a program to pass a slice of a string to a
// function (variadic) to loop through the elements and print them.

func main() {

	names := []string{"Amy", "Rob", "Helen"}
	echo(names...) // var...

}

func echo(names ...string) { // type...

	for _, s := range names {

		fmt.Print(s, " ")

	}

}
