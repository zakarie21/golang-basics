package main

import "fmt"

/*
Create a Drone struct with fields ID, Battery, and InUse (bool).

Create a slice called drones to store multiple Drone objects.

Write a function AddDrone(drones *[]Drone, id string, battery int) that adds a new drone to the list.

Write a method StartDelivery() that changes InUse to true if the battery is above 20, otherwise print "Battery too low".

Write a function ShowDrones(drones []Drone) that prints each drone’s ID, battery, and status.
*/

type Drone struct{
	ID string
	Battery int
	InUse bool
}

func addDrone(drones *[]Drone, id string, battery int){
	newDrone:= Drone{
		ID: id,
		Battery: battery,
		InUse: false,
	}

	*drones = append(*drones, newDrone)
}

func (d *Drone) startDelivery(){
	if d.Battery > 20 {
		d.InUse = true
	} else { 
		fmt.Println("Battery too low")
	}
}

func showDrones(drones []Drone) {
	
	for _, value := range drones{
		fmt.Printf("ID: %s | Battery: %d | Status: %v\n", value.ID, value.Battery, value.InUse)
	
	}
}


func main() {
	//Create empty slice of drones
	drones:= []Drone{}

	//create drones
	addDrone(&drones,"9999",45)
	addDrone(&drones, "1111",  65)
	addDrone(&drones, "2828", 15)
	addDrone(&drones,  "4774", 5 )

	fmt.Println("Inuse before")
	showDrones(drones)
	drones[3].startDelivery()
	drones[0].startDelivery()
	fmt.Println("Inuse after")
	showDrones(drones)


}