// single line comment

/*

multi line comment

*/

package main

import "fmt"

func main() {

	var i uint8 = 20
	var f float32 = 214.437
	var b bool = true
	var msg string = "Hello"

	type myStringT string
	var myMsg myStringT = "Special hello"

	fmt.Printf("%d %v %T \n", i, i, i)
	fmt.Printf("%5.1f %.4f %g %T \n", f, f, f, f)
	fmt.Printf("%t %v %T \n", b, b, b)
	fmt.Printf("%s %T \n", msg, msg)
	fmt.Printf("%v %T \n", myMsg, myMsg)

}
