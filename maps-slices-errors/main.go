package main

import "fmt"

// AddProduct: add a product if it is not already in the map
func AddProduct(inventory map[string]int, name string, stock int) {
	if inventory[name] > 0 {
		fmt.Println("Error: product already exists:", name)
	} else {
		inventory[name] = stock
		fmt.Println("Added product:", name, "with stock", stock)
	}
}

// ShowProducts: return all product names
func ShowProducts(inventory map[string]int) []string {
	products := []string{}
	for product := range inventory {
		products = append(products, product)
	}
	return products
}

// CheckStock: show stock if exists, else say not found
func CheckStock(inventory map[string]int, name string) {
	if inventory[name] == 0 {
		fmt.Println("Error: product not found:", name)
	} else {
		fmt.Println("Stock for", name, "is", inventory[name])
	}
}

func main() {
	inventory := make(map[string]int)

	inventory["Apple"] = 10
	fmt.Println("Added Apple directly with stock 10")

	AddProduct(inventory, "Banana", 5)
	AddProduct(inventory, "Apple", 3)

	fmt.Println("Products in store:", ShowProducts(inventory))

	CheckStock(inventory, "Banana")
	CheckStock(inventory, "Orange")
}
