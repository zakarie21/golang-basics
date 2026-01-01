package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Flight struct with fields FlightNumber, Destination, and SeatsAvailable
2. Make a slice called flights to store multiple flights
3. Write AddFlight(flights []Flight, f Flight) []Flight to add new flights
4. Write BookSeat(flights []Flight, flightNumber string) ([]Flight, error) to reduce seats or return an error
5. Write ShowAllFlights(flights []Flight) to display all flight details
6. FindFlightByNumber: return a flight by flight number or an error if not found
*/

type Flight struct {
	FlightNumber   string
	Destination    string
	SeatsAvailable int
}

func AddFlight(flights []Flight, f Flight) []Flight {
	flights = append(flights, f)
	return flights
}

func BookSeat(flights []Flight, flightNumber string) ([]Flight, error) {
	for i := 0; i < len(flights); i++ {
		if flights[i].FlightNumber == flightNumber {
			if flights[i].SeatsAvailable > 0 {
				flights[i].SeatsAvailable = flights[i].SeatsAvailable - 1
				return flights, nil
			} else {
				return flights, errors.New("no seats available on this flight")
			}
		}
	}
	return flights, errors.New("flight not found")
}

func ShowAllFlights(flights []Flight) {
	for _, f := range flights {
		fmt.Println("Flight:", f.FlightNumber, "| Destination:", f.Destination, "| Seats left:", f.SeatsAvailable)
	}
}

// Question 6
func FindFlightByNumber(flights []Flight, flightNumber string) (*Flight, error) {
	for i := range flights {
		if flights[i].FlightNumber == flightNumber {
			return &flights[i], nil
		}
	}
	return nil, errors.New("flight not found")
}

func main() {
	// Step 1: Create an empty list of flights
	flights := []Flight{}

	// Step 2: Add flights to the list
	flights = AddFlight(flights, Flight{FlightNumber: "GT324", Destination: "Paris", SeatsAvailable: 3})
	flights = AddFlight(flights, Flight{FlightNumber: "DC265", Destination: "New York", SeatsAvailable: 2})
	flights = AddFlight(flights, Flight{FlightNumber: "ZO999", Destination: "Berlin", SeatsAvailable: 1})

	// Step 3: Show all available flights
	fmt.Println("Available flights:")
	ShowAllFlights(flights)

	// Step 4: Try booking a seat
	var err error
	flights, err = BookSeat(flights, "VS456")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Seat booked successfully on VS456!")
	}

	// Step 5: Try booking on a non-existent flight
	flights, err = BookSeat(flights, "AA999")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 6: Show updated flight info
	fmt.Println("\nUpdated flight list:")
	ShowAllFlights(flights)

	// Demo Question 6
	flight, err := FindFlightByNumber(flights, "GT324")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFound flight:", flight.FlightNumber, "| Destination:", flight.Destination)
	}
}
