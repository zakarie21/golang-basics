package models

import "fmt"

// "fmt"


type Truck struct{
	
	BaseVehicle BaseVehicle
}

func (t *Truck) ReportStatus() string {
	return t.BaseVehicle.Status
}

func (t *Truck) MoveTo(latitude float64, longitude float64)  {
	 t.BaseVehicle.Location[0] = latitude
	 t.BaseVehicle.Location[1]= longitude
	 fmt.Println(t, "inside stuct")
}

func (t *Truck) CreateTruck(id string, location [2]float64, status string ) *Truck {
	t.BaseVehicle.ID = id
	t.BaseVehicle.Location = location
	t.BaseVehicle.Status = status

	return t
}