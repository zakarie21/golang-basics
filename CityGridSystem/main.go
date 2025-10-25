package main

import "fmt"

/*
Questions:
1. Create a map called powerGrid with city name as key and power usage as value
2. Add three cities with example power usage
3. Write a function ShowPower() to print cities and their usage
4. Write a function AddUsage() to increase power usage for a city
5. Write a function FindHighestUsage() to find the city with the highest power usage
6. Write a function RemoveCity() to remove a city from the map if it exists
7. Write a function FindCity() to check if a city exists and print its usage
8. Write a function AverageUsage() to calculate and print the average power usage
9. In main, call all functions to test the system
*/

func ShowPower(powerGrid map[string]int) {
	for city, usage := range powerGrid {
		fmt.Println(city, ":", usage, "MW")
	}
}

func AddUsage(powerGrid map[string]int, city string, amount int) {
	powerGrid[city] = powerGrid[city] + amount
}

func FindHighestUsage(powerGrid map[string]int) {
	highestCity := ""
	highestUsage := 0

	for city, usage := range powerGrid {
		if usage > highestUsage {
			highestUsage = usage
			highestCity = city
		}
	}

	fmt.Println("\nCity with highest power usage:", highestCity, "-", highestUsage, "MW")
}

func RemoveCity(powerGrid map[string]int, city string) {
	if _, exists := powerGrid[city]; exists {
		delete(powerGrid, city)
		fmt.Println("\nRemoved city:", city)
	} else {
		fmt.Println("\nCity not found:", city)
	}
}

func FindCity(powerGrid map[string]int, city string) {
	if usage, exists := powerGrid[city]; exists {
		fmt.Println("\nCity found:", city, "-", usage, "MW")
	} else {
		fmt.Println("\nCity not found:", city)
	}
}

func AverageUsage(powerGrid map[string]int) {
	if len(powerGrid) == 0 {
		fmt.Println("\nNo cities in the power grid.")
		return
	}

	total := 0
	for _, usage := range powerGrid {
		total = total + usage
	}

	average := total / len(powerGrid)
	fmt.Println("\nAverage power usage across all cities:", average, "MW")
}

func main() {
	// Step 1: Create a map for city power usage
	powerGrid := map[string]int{
		"London":      900,
		"Manchester":  750,
		"Birmingham":  820,
		"Liverpool":   700,
		"Leeds":       680,
		"Glasgow":     770,
	}

	// Step 2: Show starting power usage
	fmt.Println("Initial power usage:")
	ShowPower(powerGrid)

	// Step 3: Increase London’s power usage
	AddUsage(powerGrid, "London", 150)

	// Step 4: Show updated power usage
	fmt.Println("\nAfter adding usage:")
	ShowPower(powerGrid)

	// Step 5: Find and show the city with the highest usage
	FindHighestUsage(powerGrid)

	// Step 6: Remove a city
	RemoveCity(powerGrid, "Leeds")

	// Step 7: Try finding a city
	FindCity(powerGrid, "London")
	FindCity(powerGrid, "Leeds")

	// Step 8: Show average power usage
	AverageUsage(powerGrid)

	// Step 9: Show final map
	fmt.Println("\nFinal power grid:")
	ShowPower(powerGrid)
}

