package models

//import "fmt"


type Circle struct{
	Radius float64 `json:"radius"`
}

func (c Circle) Perimeter() float64{
	return 2*3.14*c.Radius
}