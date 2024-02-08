package main

import (
	"fmt"
)

func main() {
	// Prompt the user to input scores for three subjects
	var mathScore, scienceScore, englishScore float64
	fmt.Println("Enter your scores for Math, Science, and English:")

	// Take input for each subject's score
	fmt.Print("Math score: ")
	fmt.Scanln(&mathScore)
	fmt.Print("Science score: ")
	fmt.Scanln(&scienceScore)
	fmt.Print("English score: ")
	fmt.Scanln(&englishScore)

	// Calculate the average score
	averageScore := (mathScore + scienceScore + englishScore) / 3

	// Determine the corresponding letter grade based on the average score
	var letterGrade string
	switch {
	case averageScore >= 90:
		letterGrade = "A"
	case averageScore >= 80:
		letterGrade = "B"
	case averageScore >= 70:
		letterGrade = "C"
	case averageScore >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	// Display the average score and corresponding letter grade
	fmt.Printf("\nAverage Score: %.2f\n", averageScore)
	fmt.Println("Letter Grade:", letterGrade)
}
