package main

import "fmt"

var i int = 10

func main() {

	f1()

}

func f1() {

	j := 20
	k := 30
	fmt.Println("f1(#1) - ", i, j, k)
	f2(j)
	fmt.Println("f1(#2) - ", i, j, k)

}

func f2(jp int) {

	fmt.Println("f2(#1) - ", i, jp)
	i = 11
	i := 30

	f3(i, jp)
	fmt.Println("f1(#2) - ", i, jp)

}

func f3(ip, jpp, int) {

	ip++
	jpp++
	fmt.Println("f3(#1) - ", i, ip, jpp)

}
