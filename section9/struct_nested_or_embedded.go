package main

import "fmt"

type generalInfo struct {
	country, hairColour string
}

type player struct {
	name, sport string
	age         int
	info        generalInfo
}

func main() {

	var player1 player
	player1.name = "Wayne Gretzky"
	player1.age = 57
	player1.sport = "Hockey"
	player1.info.country = "Canada"
	player1.info.hairColour = "Brown"

	fmt.Println("Player1 =", player1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	info2 := generalInfo{
		country:    "USA",
		hairColour: "black",
	}

	player2 := player{
		name:  "Lebron James",
		sport: "basketball",
		age:   32,
		info:  info2,
	}

	fmt.Println("Player2 =", player2)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player3 := player{
		name:  "Serena Williams",
		sport: "Tennis",
		age:   36,
		info: generalInfo{
			country:    "USA",
			hairColour: "black",
		},
	}

	fmt.Println("Player3 =", player3)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player4 := player{"Anderson Silva", "MMA", 43, generalInfo{"Brazil", "Black"}}
	fmt.Println("Player4 =", player4)

	p4 := &player4
	fmt.Println("*p4:", *p4)
	fmt.Printf("(*p4).name=%s (*p4).info.country=%s", (*p4).name, (*p4).info.country)

}
