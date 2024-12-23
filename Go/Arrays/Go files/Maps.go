package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to create and populate the map
func makeMap(lenMap int) map[string]int {
	mapp := make(map[string]int) // Create a map with the given capacity

	for i := 0; i < lenMap; i++ {
		var name string
		var ageInput string

		// Get the name for the entry
		fmt.Printf("Enter name for entry %d: ", i+1)
		fmt.Scan(&name)
		name = strings.TrimSpace(name) // Remove trailing spaces

		// Get the age for the name
		fmt.Printf("Enter age for %s: ", name)
		fmt.Scan(&ageInput)
		ageInput = strings.TrimSpace(ageInput)

		// Convert the age input to integer
		age, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("Invalid age input. Please enter a valid number.")
			i-- // Retry the same entry
			continue
		}

		mapp[name] = age
	}
	return mapp
}

func main() {
	var exit string
	var lenMap int

	// Read map length from user
	fmt.Print("What should be the length of the map: ")
	_, err := fmt.Scan(&lenMap)
	if err != nil || lenMap < 1 {
		fmt.Println("Invalid input for map length.")
		fmt.Print("Press Enter to exit... ")
		fmt.Scan(&exit)
		return
	}

	// Create and populate the map
	mapp := makeMap(lenMap)

	// Display map contents
	fmt.Println("\nContents of the map:")
	for name, age := range mapp {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// Display map length
	fmt.Printf("\nThis is the length of the map: %d\n", len(mapp))

	// Wait for user input to exit
	fmt.Print("\nPress Enter to exit... ")
	fmt.Scan(&exit)
}

