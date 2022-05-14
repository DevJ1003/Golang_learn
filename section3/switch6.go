package main

import "fmt"

func main() {

	sportID := 2
	isYoung := true

	switch {

	case sportID == 1 && isYoung:

		fmt.Println("Playing soccer...")

	case sportID == 1 && !isYoung:

		fmt.Println("Watching soccer...")

	case sportID == 2 && isYoung:

		fmt.Println("Playing chess...")

	case sportID == 2 && !isYoung:

		fmt.Println("Still can play chess...")

	default:

		fmt.Println("Cannot play any game rather than soccer and chess...")

	}

}
