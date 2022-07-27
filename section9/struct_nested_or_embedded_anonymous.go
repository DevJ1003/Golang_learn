package main

import "fmt"

// ASSIGNMENT : embedded anonymous structs

type generalInfo struct {
	country, hairColour string
}

type player struct {
	name, sport string
	age         int
	generalInfo // anonymous
}

func main() {

	var player1 player
	player1.name = "Wayne Gretzky"
	player1.age = 57
	player1.sport = "Hockey"
	player1.country = "Canada"
	player1.hairColour = "Brown"

	fmt.Println("Player1 =", player1)

}
