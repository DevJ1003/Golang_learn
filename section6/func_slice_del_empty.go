package main

import "fmt"

// TO-DO : write a program to accept a slice of
// string and return a new slice with no blank strings.
// Use both variadic and regular function.

func main() {

	data := []string{"Daisy", "Rose", " ", "Tulip"}

	fmt.Printf("%q\n", trimSlice(data)) // form1
	// fmt.Printf("%q\n", trimSlice(data...)) // form2

	fmt.Printf("%q\n", data)

}

func trimSlice(data []string) []string { // form1
	// func trimSlice(data ...string) []string{ // form2

	newData := data[:0] // note
	// var newData []string

	for _, d := range data {

		if d != " " {

			newData = append(newData, d)

		}

	}

	return newData

}
