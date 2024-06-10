package main

import (
	"fmt"
	"strconv"
)

type Subject struct {
	name  string
	units float64
	grade float64
}

func main() {
	var subjects [3]Subject

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter details for subject %d:\n", i+1)

		fmt.Print("Subject name: ")
		fmt.Scanln(&subjects[i].name)

		fmt.Print("Number of units: ")
		var unitsStr string
		fmt.Scanln(&unitsStr)
		units, err := strconv.ParseFloat(unitsStr, 64)
		if err != nil {
			fmt.Println("Invalid input for units. Please enter a number.")
			i--
			continue
		}
		subjects[i].units = units

		fmt.Print("Grade: ")
		var gradeStr string
		fmt.Scanln(&gradeStr)
		grade, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil {
			fmt.Println("Invalid input for grade. Please enter a number.")
			i--
			continue
		}
		subjects[i].grade = grade
	}

	gwa := calculateGWA(subjects[:])
	fmt.Printf("Your GWA is: %.2f\n", gwa)
}

func calculateGWA(subjects []Subject) float64 {
	var totalGradeUnits, totalUnits float64
	for _, subject := range subjects {
		totalGradeUnits += subject.grade * subject.units
		totalUnits += subject.units
	}
	return totalGradeUnits / totalUnits
}