package main

import "fmt"

func main() {

	values := []string{"AB"}

	fmt.Println(values)

	for v := range values {

		// bytes := []byte(values[v])

		fmt.Println(v)

		bytes := []byte("AB")

		fmt.Println(bytes)

	}

}
