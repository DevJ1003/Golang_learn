package main

import (
	"fmt"
	"time"
)

func main() {

	var DelayTime = 500 * time.Millisecond

	fmt.Println("Press CTRL + C to exit the program!")
	fmt.Print("Operation in progress ...")

	for {

		fmt.Print(`\`)
		time.Sleep(DelayTime)
		fmt.Print("b")

		fmt.Print(`|`)
		time.Sleep(DelayTime)
		fmt.Print("b")

		fmt.Print(`/`)
		time.Sleep(DelayTime)
		fmt.Print("b")

		fmt.Print(`-`)
		time.Sleep(DelayTime)
		fmt.Print("b")

	}

}
