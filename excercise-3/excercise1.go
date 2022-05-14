package main

import "fmt"

func main() {

	const KilogramsPerPound = 0.453592
	const MetersInInch = 0.0254

	var weight float64
	fmt.Print("\nEnter weight (in pounds) : ")
	fmt.Scanf("%f ", &weight)

	var height float64
	fmt.Print("Enter height (in inches) : ")
	fmt.Scanf("%f ", &height)

	// Converting values into koilgrams and meters from pounds and inches
	WeightInKilograms := weight * KilogramsPerPound
	HeightInMeters := height * MetersInInch

	// Calculation of BMI
	BMI := WeightInKilograms / (HeightInMeters * HeightInMeters)

	WeightStatus := ""

	if BMI < 18.5 {

		WeightStatus = "Underweight"

	} else if BMI < 25 {

		WeightStatus = "Normal"

	} else if BMI < 30 {

		WeightStatus = "Overweight"

	} else {

		WeightStatus = "Obese"

	}

	fmt.Println()
	fmt.Printf("%s %10s %10s %10s\n", "Weight", "Height", "BMI", "Status")
	fmt.Printf("\n%.2f %10.2f %11.2f %14s", weight, height, BMI, WeightStatus)

}
