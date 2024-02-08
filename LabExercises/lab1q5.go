package main

import (
	"fmt"
)

func main() {
	var numSets int
	fmt.Println("Enter the number of sets of data:")
	fmt.Scanln(&numSets)

	principalAmounts := make([]float64, numSets)
	annualInterestRates := make([]float64, numSets)
	years := make([]int, numSets)

	for i := 0; i < numSets; i++ {
		fmt.Printf("\nEnter details for set %d:\n", i+1)
		fmt.Print("Principal Amount: ")
		fmt.Scanln(&principalAmounts[i])
		fmt.Print("Annual Interest Rate (%): ")
		fmt.Scanln(&annualInterestRates[i])
		fmt.Print("Number of Years: ")
		fmt.Scanln(&years[i])
	}

	for i := 0; i < numSets; i++ {
		interest := (principalAmounts[i] * annualInterestRates[i] * float64(years[i])) / 100
		fmt.Printf("\nFor set %d:\n", i+1)
		fmt.Printf("Principal Amount: $%.2f\n", principalAmounts[i])
		fmt.Printf("Annual Interest Rate: %.2f%%\n", annualInterestRates[i])
		fmt.Printf("Number of Years: %d\n", years[i])
		fmt.Printf("Simple Interest: $%.2f\n", interest)
	}
}
