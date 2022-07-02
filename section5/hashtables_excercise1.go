package main

import (
	"fmt"
	"math"
)

func main() {

	values := []string{"tea", "tea", "eat", "eta", "teal", "teatop", "teaparty", "tear"}

	for v := range values {

		hashkey := hashK(values[v])
		fmt.Printf("%-10s (hashkey : %10d) %v\n", values[v], hashkey, []byte(values[v]))

	}

}

func hashK(s string) int {

	result := 0
	idx := 0

	for i := len(s) - 1; i > -1; i-- {

		result += int(math.Pow10(i)) * int(s[idx])
		idx++

	}

	return result

}
