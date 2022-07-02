package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "你好吗 你最近怎么样 你的年龄 你学什么"

	r := []rune(s)

	// fmt.Println(r)
	fmt.Println(len(s), utf8.RuneCountInString(s))

	counts := make(map[rune]int)

	fmt.Println()
	for i := range r {

		counts[r[i]]++

	}

	fmt.Printf("%v\n", counts)

	for k, v := range counts {

		fmt.Printf("%7d %3s %3d \n", k, string(k), v)

	}

}
