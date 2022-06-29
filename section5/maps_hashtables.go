package main

import "fmt"

func main() {

	values := []string{"ABC", "BCA", "CAB", "CBA", "ACB", "BAC"}
	factor := []int{100, 10, 1}

	hashkey := 0
	for v := range values {

		bytes := []byte(values[v])
		f := 0
		hashkey = 0
		for i := range bytes {

			fmt.Print(bytes[i], " ")
			hashkey += int(bytes[i]) * factor[f]
			f++

		}

		fmt.Printf(" (hashkey : %d) \n", hashkey)

	}

}
