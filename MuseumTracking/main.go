package main

import (
	"fmt"
	
)

/*
Create an Artifact struct with fields Name, Age (int), and OnDisplay (bool).

Write a function AddArtifact(collection *[]Artifact, name string, age int) that adds a new artifact to the museum’s collection.

Write a method Display() that sets OnDisplay to true if the artifact’s age is above 50, otherwise prints "Too new to display".

Write a function ShowCollection(collection []Artifact) that prints all artifacts and their display status.

Write a function FindArtifact(collection []Artifact, name string) that searches by name and returns an error if not found.
*/

type Artifact struct {
	Name string
	Age int
	OnDisplay bool
}


func addArtifact(collection *[]Artifact, name string, age int) {
	newCollection:=Artifact{
		Name : name,
		Age: age,
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
		fmt.Printf("Name: %s | Age: %d | OnDisplay: %v \n", value.Name, value.Age,value.OnDisplay)
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
	


}



