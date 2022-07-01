package main

import "fmt"

func main() {

	values := []string{"AB"}

	fmt.Println(values)

	for v := range values {

		fmt.Println("v's range : ", v)
		// fmt.Println(v)

		bytes := []byte(values[v])

		fmt.Println(bytes)

		for i := range bytes {

			fmt.Println(i)
			fmt.Printf("\n")
			// fmt.Print(bytes[i], " ")

		}

	}

}
