package models

import (
	// "fmt"
)

type Drone struct{
	
	BaseVehicle BaseVehicle
}



func (d Drone) ReportStatus() string {
	return d.BaseVehicle.Status
}

func (t Drone) MoveTo(latitude float64, longitude float64)  {
	 t.BaseVehicle.Location[0] = latitude
	 t.BaseVehicle.Location[1]= longitude
}

func (d Drone) CreateDrone(id string, location [2]float64, status string ) Drone {
	d.BaseVehicle.ID = id
	d.BaseVehicle.Location = location
	d.BaseVehicle.Status = status

	return d
}