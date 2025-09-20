package main

import "fmt"

// Homework:
// 1. Create new struct vehicle with some values (engineStatus = string)
// 2. Create method fixEngine which will update the status
// 3. Create a function which changes one engine of one vehicle to another (kind of like a swap)
// 4. Create a function which finds the best vehicle in a list of vehicles

// Homework (20/09):
// 1. Make vehicles have a speed
// 2. Create a method which allows two vehicles to race and output which vehicle wins 
// 2.1. (if engine is broken, that vehicle will lose)
// 2.2. each race will run some miles on the vehicle
// 2.3. we should be able to specify how long a race is (miles)


type Vehicle struct {
	Brand        string
	Year         int
	EngineStatus string
	Miles        int
    Speed    int
}

func (v Vehicle) carRace(car2 Vehicle) string{
    if v.EngineStatus == "Broken"{
        return car2.Brand
    } else if car2.EngineStatus == "Broken"{
        return v.Brand
    }

    if v.Speed < car2.Speed{
        return car2.Brand
    } else {
		return  v.Brand
	}
}

func deductMiles(car1, car2 *Vehicle, miles int) {
	car1.Miles = car1.Miles + miles
	car2.Miles = car2.Miles + miles
}

func main() {
   

	car1 := Vehicle{"Toyota", 2025, "Broken", 20000, 250}
	car2 := Vehicle{"Ford", 2024, "Fixed", 30000,300}

	raceResult:= car1.carRace(car2)
	fmt.Println("The faster car is", raceResult)

	fmt.Println("Before race miles on car 1:", car1.Miles)
	fmt.Println("Before race miles on car 2:", car2.Miles)
	deductMiles(&car1, &car2, 7000)
	fmt.Println("After race miles on car 1:", car1.Miles)
	fmt.Println("After race miles on car 2:", car2.Miles)

	fmt.Println(car1.EngineStatus)
	car1.fixEngine()
	fmt.Println(car1.EngineStatus)

	changeEngine(&car1, &car2)
	fmt.Println("New engine status for Car 1 is", car1.EngineStatus)
	fmt.Println("New engine status for Car 2 is", car2.EngineStatus)

	// testNumbers := []int{2, 3, 1}
	// resultN := findHighestValue(testNumbers)

	testCars := []Vehicle{car2,car2}

	result := findBestCar(testCars)
	fmt.Printf("The best car is %s as it is the most recent\n", result.Brand)

}


func (v *Vehicle) fixEngine() {

	if v.EngineStatus != "Fixed" {
		v.EngineStatus = "Fixed"
		fmt.Println("New engine status is fixed")
	} else {
		fmt.Println("Engine status is already", v.EngineStatus)
	}

}

func changeEngine(vehicleOne, vehicleTwo *Vehicle) {
	vehicleThree := vehicleOne.EngineStatus
	vehicleOne.EngineStatus = vehicleTwo.EngineStatus
	vehicleTwo.EngineStatus = vehicleThree

}

func findBestCar(vehicles []Vehicle) Vehicle {

	myArray := [3]int{}
	fmt.Println(myArray[2])

	bestCar := vehicles[0]
	for _, v := range vehicles {
		if v.Year > bestCar.Year {
			bestCar = v
		}
	}

	return bestCar
}


// Homework:
// 1. Create new struct vehicle with some values (engineStatus = string)
// 2. Create method fixEngine which will update the status
// 3. Create a function which changes one engine of one vehicle to another (kind of like a swap)
// 4. Create a function which finds the best vehicle in a list of vehicles
