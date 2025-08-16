package models

//import "fmt"

type Rectangle struct{
	Width float64  `json:"Width"`
	Height float64  `json:"Height"`
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.Height+r.Width)
}