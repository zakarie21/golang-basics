package main

import "fmt"

func main() {

    // 1. Define a struct called Person with name and age
    // Answer: A struct groups related data together
    type Person struct {
        name string
        age  int
    }

    // Create an instance of Person
    person1 := Person{name: "Alice", age: 25}
    fmt.Println("1:", person1.name, person1.age)

    // 2. Access and modify fields
    // Answer: You can read and update fields using dot notation
    person1.age = 26 // change age
    fmt.Println("2: Updated age:", person1.age)
    fmt.Println("2: Updated struct:", person1.name, person1.age)

    // 3. Struct as function parameter
    // Answer: You can pass the whole struct to a function
    printGreeting := func(p Person) {
        message := "Hello, " + p.name + "!"
        fmt.Println("3:", message)
    }
    printGreeting(person1)

    // 4. Struct slice
    // Answer: You can make a slice of structs and loop through them
    person2 := Person{name: "Bob", age: 30}
    person3 := Person{name: "Charlie", age: 22}
    people := []Person{person1, person2, person3}

    fmt.Println("4: List of people")
    for _, p := range people {
        fmt.Println(p.name, "-", p.age)
    }
}
