package main

import (
	"fmt"
)

/*
Questions:
1. Create an Artifact struct with fields Name, Age (int), and OnDisplay (bool).
2. Write a function AddArtifact(collection *[]Artifact, name string, age int) that adds a new artifact to the museum’s collection.
3. Write a method Display() that sets OnDisplay to true if the artifact’s age is above 50, otherwise prints "Too new to display".
4. Write a function ShowCollection(collection []Artifact) that prints all artifacts and their display status.
5. Write a function FindArtifact(collection []Artifact, name string) that searches by name and returns an error if not found.
6. Write a function CountOnDisplay(collection []Artifact) int that returns the number of artifacts currently on display.
7. Write a function OldestArtifact(collection []Artifact) Artifact that returns the artifact with the highest age.
8. Write a function CountNewArtifacts(collection []Artifact, limit int) int that returns how many artifacts are newer than the given age limit.
9. Write a function AverageAge(collection []Artifact) float64 that returns the average age of all artifacts.
*/

type Artifact struct {
	Name      string
	Age       int
	OnDisplay bool
}

func addArtifact(collection *[]Artifact, name string, age int) {
	newCollection := Artifact{
		Name: name,
		Age:  age,
	}
	*collection = append(*collection, newCollection)
}

func (a *Artifact) Display() {
	if a.Age > 50 {
		a.OnDisplay = true
	} else {
		fmt.Printf("%s is too new to display\n", a.Name)
	}
}

func showCollection(collection []Artifact) {
	for _, value := range collection {
		fmt.Printf("Name: %s | Age: %d | OnDisplay: %v \n", value.Name, value.Age, value.OnDisplay)
	}
}

func FindArtifact(collection []Artifact, name string) {
	found := false
	for _, value := range collection {
		if value.Name == name {
			fmt.Printf("Artifact %s was found\n", value.Name)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Artifact %s could not be found\n", name)
	}
}

func CountOnDisplay(collection []Artifact) int {
	count := 0
	for _, item := range collection {
		if item.OnDisplay {
			count++
		}
	}
	return count
}

func OldestArtifact(collection []Artifact) Artifact {
	if len(collection) == 0 {
		return Artifact{}
	}

	oldest := collection[0]
	for _, item := range collection {
		if item.Age > oldest.Age {
			oldest = item
		}
	}
	return oldest
}

func CountNewArtifacts(collection []Artifact, limit int) int {
	count := 0
	for _, item := range collection {
		if item.Age < limit {
			count++
		}
	}
	return count
}


func AverageAge(collection []Artifact) float64 {
	if len(collection) == 0 {
		return 0
	}
	total := 0
	for _, item := range collection {
		total += item.Age
	}
	return float64(total) / float64(len(collection))
}

func main() {

	collection := []Artifact{}

	addArtifact(&collection, "MonaLisa", 150)
	addArtifact(&collection, "Bible", 2150)
	addArtifact(&collection, "Crown", 30)
	addArtifact(&collection, "Ship", 150)

	fmt.Println("Collection before:")
	showCollection(collection)

	collection[0].Display()
	collection[1].Display()
	collection[3].Display()

	fmt.Println("Collection after:")
	showCollection(collection)

	fmt.Println("")
	collection[2].Display()
	showCollection(collection)

	fmt.Println("")
	FindArtifact(collection, "Bible")
	FindArtifact(collection, "Omar")

	fmt.Println("\nNumber of artifacts on display:", CountOnDisplay(collection))

	oldest := OldestArtifact(collection)
	fmt.Printf("\nOldest artifact: %s (Age: %d)\n", oldest.Name, oldest.Age)

	
	fmt.Printf("\nAverage artifact age: %.2f years\n", AverageAge(collection))
}
