package main


import (
	"fmt"
	models "vehiclemNgementsystem/models"
)

type Vehicle interface {
	MoveTo (latitude, longitude float64)
	ReportStatus() string
}

func MovingVehicle(v Vehicle, latitude, longitude float64) Vehicle{
	v.MoveTo(latitude,longitude)
	
	return v
}

func main(){

	newTruckPtr := &models.Truck{}
	
	newTruck := newTruckPtr.CreateTruck("45", [2]float64{2,4}, "inTransit")
	// newDrone := models.Drone{}.CreateDrone("55", [2]float64{12,24}, "inTransit")

	
	fmt.Println(MovingVehicle(newTruck, 45,55))

}