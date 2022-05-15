package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const MinRank = 1
	const MaxRank = 13

	const MinSuit = 1
	const MaxSuit = 4

	rand.Seed(time.Now().UnixNano())

	rankIdx := MinRank + rand.Intn(MaxRank-MinRank)

	suitIdx := MinSuit + rand.Intn(MaxSuit-MinSuit)

	rankString := ""
	switch rankIdx {

	case 1:
		rankString = "Ace"

	case 2:
		rankString = "2"

	case 3:
		rankString = "3"

	case 4:
		rankString = "4"

	case 5:
		rankString = "5"

	case 6:
		rankString = "6"

	case 7:
		rankString = "7"

	case 8:
		rankString = "8"

	case 9:
		rankString = "9"

	case 10:
		rankString = "10"

	case 11:
		rankString = "Jack"

	case 12:
		rankString = "Queen"

	case 13:
		rankString = "King"

	}

	suitString := ""

	switch suitIdx {

	case 1:
		suitString = "Clubs"

	case 2:
		suitString = "Diamonds"

	case 3:
		suitString = "Hearts"

	case 4:
		suitString = "Spades"

	}

	fmt.Printf("You picked %s of %s.\n", rankString, suitString)

}
