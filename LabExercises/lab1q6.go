package main

import (
	"fmt"
)

const numGameTitles = 5 // Number of different game titles available

var gameTitles = [numGameTitles]string{"Title 1", "Title 2", "Title 3", "Title 4", "Title 5"}

func main() {
	var gameQuantities [numGameTitles]int // Quantities available for each game title
	var choice int

	// Input quantities for each game title
	for i, title := range gameTitles {
		fmt.Printf("Enter quantity available for %s: ", title)
		fmt.Scanln(&gameQuantities[i])
	}

	// Menu for inventory management system
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. View Inventory")
		fmt.Println("2. Update Inventory")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// View inventory
			fmt.Println("\nInventory:")
			for i, title := range gameTitles {
				fmt.Printf("%s: %d\n", title, gameQuantities[i])
			}
		case 2:
			// Update inventory
			var index int
			fmt.Print("Enter index of the game title to update: ")
			fmt.Scanln(&index)
			if index >= 0 && index < numGameTitles {
				fmt.Printf("Enter new quantity for %s: ", gameTitles[index])
				fmt.Scanln(&gameQuantities[index])
				fmt.Println("Inventory updated successfully!")
			} else {
				fmt.Println("Invalid index!")
			}
		case 3:
			// Exit the program
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
