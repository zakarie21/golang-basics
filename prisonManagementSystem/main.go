package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Prisoner struct with fields Name, CellNumber, and SentenceYears
2. Make a slice called prisoners to store multiple prisoners
3. Write AddPrisoner function to add a new prisoner
4. Write ReleasePrisoner function to remove a prisoner by name or return an error if not found
5. Write CountYearsLeft function that returns a map showing each prisoner's name and their remaining sentence years
*/

type Prisoner struct {
	Name          string
	CellNumber    int
	SentenceYears int
}

func AddPrisoner(prisoners []Prisoner, p Prisoner) []Prisoner {
	prisoners = append(prisoners, p)
	return prisoners
}

func ReleasePrisoner(prisoners []Prisoner, name string) ([]Prisoner, error) {
	for i := 0; i < len(prisoners); i++ {
		if prisoners[i].Name == name {
			prisoners = append(prisoners[:i], prisoners[i+1:]...)
			return prisoners, nil
		}
	}
	return prisoners, errors.New("prisoner not found")
}

func CountYearsLeft(prisoners []Prisoner) map[string]int {
	result := make(map[string]int)
	for _, p := range prisoners {
		result[p.Name] = p.SentenceYears
	}
	return result
}

func main() {
	// Step 1: Create an empty slice of prisoners
	prisoners := []Prisoner{}

	// Step 2: Add prisoners to the list
	prisoners = AddPrisoner(prisoners, Prisoner{Name: "Zak Omar", CellNumber: 101, SentenceYears: 5})
	prisoners = AddPrisoner(prisoners, Prisoner{Name: "Tom Peter", CellNumber: 102, SentenceYears: 8})
	prisoners = AddPrisoner(prisoners, Prisoner{Name: "Bruce Wayne", CellNumber: 103, SentenceYears: 3})

	// Step 3: Try to release one prisoner
	var err error
	prisoners, err = ReleasePrisoner(prisoners, "Jake Brown")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Jake Brown has been released.")
	}

	// Step 4: Print remaining prisoners
	fmt.Println("\nCurrent Prisoners:")
	for _, p := range prisoners {
		fmt.Println(p.Name, "in Cell", p.CellNumber, "with", p.SentenceYears, "years left.")
	}

	// Step 5: Count years left for all prisoners
	fmt.Println("\nYears Left by Prisoner:")
	yearsLeft := CountYearsLeft(prisoners)
	for name, years := range yearsLeft {
		fmt.Println(name, "has", years, "years remaining.")
	}
}
