package main

import (
	"errors"
	"fmt"
)

// 1. Create Animal and Zoo structs
// 2. Add AddAnimal method that puts an animal into the zoo
// 3. Create Feedable interface and implement Feed method for Animal
// 4. Add Release method to Animal with error handling
// 5. Add CountAnimals method to Zoo that returns how many animals are inside

type Animal struct {
	Name      string
	Species   string
	Available bool
}

func (a *Animal) Feed() error {
	if a.Available == false {
		return errors.New("animal already fed or not available")
	}
	a.Available = false
	return nil
}

func (a *Animal) Release() error {
	if a.Available == false {
		return errors.New("animal already released")
	}
	a.Available = false
	return nil
}

type Feedable interface {
	Feed() error
}

type Zoo struct {
	Animals []Animal
}

func NewZoo() Zoo {
	return Zoo{
		Animals: []Animal{},
	}
}

func (z *Zoo) AddAnimal(a Animal) {
	z.Animals = append(z.Animals, a)
}

func (z *Zoo) CountAnimals() int {
	return len(z.Animals)
}

func main() {
	zoo := NewZoo()

	lion := Animal{Name: "Simba", Species: "Lion", Available: true}
	elephant := Animal{Name: "Dumbo", Species: "Elephant", Available: true}

	zoo.AddAnimal(lion)
	zoo.AddAnimal(elephant)

	fmt.Println("Number of animals in the zoo:", zoo.CountAnimals())

	err := zoo.Animals[0].Feed()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = zoo.Animals[0].Feed()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = zoo.Animals[0].Release()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
