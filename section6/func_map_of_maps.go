package main

import "fmt"

func main() {

	employees := map[string]map[string]string{

		"BT": map[string]string{

			"firstname": "Blake",
			"lastname":  "Travis",
		},

		"PC": map[string]string{

			"firstname": "Parker",
			"lastname":  "Cooper",
		},

		"DC": map[string]string{

			"firstname": "Dakota",
			"lastname":  "Carington",
		},
	}

	if emp, ok := employees["PC"]; ok {

		fmt.Println(emp["firstname"], emp["lastname"])

	}

	for initials, emp := range employees {

		fmt.Println(initials, emp["firstname"], emp["lastname"])

	}

}
