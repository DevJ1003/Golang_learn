package main

import "fmt"

func main() {

	player1 := struct {
		name, sport string
		age         int
	}{"Sydney Crosby", "Hoceky", 30} // positional approach

	fmt.Println("Player 1 =", player1)
	// Now, above struct cannot be used again.

	player2 := struct {
		name, sport string
		age         int
	}{

		age:   21,
		sport: "Swimming",
		name:  "Katie Ledecky",
	}

	fmt.Println("Player 2 =", player2)

}
