package main

import (
	"fmt"
)

type Laptop struct {
	Brand   string
	Model   string
	Year    int
	Battery int
}

func (l *Laptop) Charge(amount int) {
	var newBattery int
	newBattery = l.Battery + amount
	if newBattery > 100 {
		newBattery = 100
	}
	l.Battery = newBattery
}

func UpgradeYear(l *Laptop, newYear int) {
	l.Year = newYear
}

func SwapModels(l1, l2 *Laptop) {
	var temp string
	temp = l1.Model
	l1.Model = l2.Model
	l2.Model = temp
}

func FindNewestLaptop(laptops []Laptop) Laptop {
	var newest Laptop
	newest = laptops[0]
	for _, l := range laptops {
		if l.Year > newest.Year {
			newest = l
		}
	}
	return newest
}

func main() {
	var l1, l2, l3 Laptop

	l1 = Laptop{"Dell", "XPS 13", 2020, 60}
	l2 = Laptop{"Apple", "MacBook Pro", 2021, 80}
	l3 = Laptop{"Lenovo", "ThinkPad X1", 2019, 50}

	l1.Charge(30)
	fmt.Printf("This is the charge of %q: %d%%\n", l1.Model, l1.Battery)

	UpgradeYear(&l3, 2022)
	fmt.Printf("This is the updated year of %q: %d\n", l3.Model, l3.Year)

	fmt.Printf("This is the model before swap: %q and %q\n", l1.Model, l2.Model)
	SwapModels(&l1, &l2)
	fmt.Printf("This is the model after swap: %q and %q\n", l1.Model, l2.Model)

	laptops := []Laptop{l1, l2, l3}
	newest := FindNewestLaptop(laptops)
	fmt.Printf("Final newest laptop is: %q, Brand: %q, Year: %d\n", newest.Model, newest.Brand, newest.Year)
}

