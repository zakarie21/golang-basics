package main

import "fmt"

/*
Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

*/


type Product struct{
	Title string
	ID string
	Price float64
}




func main(){

	products:= []Product{
		{Title: "Chair", ID: "345", Price: 12.50},
		{Title: "Table", ID: "123", Price: 20.75},
	}
	fmt.Println(products)
	
	products = append(products, Product{Title: "Door", ID: "999", Price: 30.12})

	fmt.Println("New prodyct range:", products)

	

	hobbies:= [3]string{"Football", "Basketball","Cricket"}
	for _, hobby := range hobbies{
		fmt.Println(hobby)
	}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	newHobbies := hobbies[1:3]
	fmt.Println(newHobbies)

	firstSlice:= hobbies[0:2]
	secondSlice := hobbies[:2]
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	 
	reslice := hobbies[1:3]
	fmt.Println(reslice)

	courseGoals:= []string{"Golang", "Java"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Python"
	courseGoals = append(courseGoals, "Typescript")
	fmt.Println(courseGoals)

	for _, course := range courseGoals{
		fmt.Println(course)
	}

	//maps for key string and value float64

	newMap:= make(map[string]float64)
	newMap["Zak"] = 40.5
	newMap["Lucy"] = 32.4
	fmt.Println(newMap)

	for i, maps := range newMap{
		fmt.Printf("Key: %s, Value: %.2f\n", i, maps)
	}

	
}