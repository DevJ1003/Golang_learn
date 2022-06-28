package main

import "fmt"

func main() {

	team1 := []string{"Messi", "Pele", "Maradona", goalkeeper()}

	team := append([]string{"Maldini", "Cafu", "Carlos", "Backenbauer"}, team1...)
	team = append(team, midfielders()...)

	for i, name := range team {

		fmt.Println(i, name)

	}

}

// FUNCTIONS
func midfielders() []string {

	return []string{"Iniesta", "Zidane", "Platini"}

}

func goalkeeper() string {

	return "Bufon"

}
