package main

import "fmt"

func main() {

	values := []string{"DEV"}

	fmt.Println(values)

	for v := range values {

		fmt.Println(v)

		bytes := []byte(values[v])

		fmt.Println(bytes)

	}

}
