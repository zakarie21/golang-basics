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
7. CountTotalFlights: return the total number of flights
8. ListAvailableFlights: return all flights that still have available seats
9. HasAvailableFlights: return true if at least one flight has available seats
10. GetFirstAvailableFlight: return the first flight with available seats or an error if none
11. CountAvailableSeats: return total available seats across all flights
12. GetFlightWithMostAvailableSeats: return flight with most seats or an error if none
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

// Question 7
func CountTotalFlights(flights []Flight) int {
	return len(flights)
}

// Question 8
func ListAvailableFlights(flights []Flight) []Flight {
	var available []Flight
	for _, flight := range flights {
		if flight.SeatsAvailable > 0 {
			available = append(available, flight)
		}
	}
	return available
}

// Question 9
func HasAvailableFlights(flights []Flight) bool {
	for _, flight := range flights {
		if flight.SeatsAvailable > 0 {
			return true
		}
	}
	return false
}

// Question 10
func GetFirstAvailableFlight(flights []Flight) (*Flight, error) {
	for i := range flights {
		if flights[i].SeatsAvailable > 0 {
			return &flights[i], nil
		}
	}
	return nil, errors.New("no available flights")
}

// Question 11
func CountAvailableSeats(flights []Flight) int {
	total := 0
	for _, flight := range flights {
		total += flight.SeatsAvailable
	}
	return total
}

// Question 12
func GetFlightWithMostAvailableSeats(flights []Flight) (*Flight, error) {
	if len(flights) == 0 {
		return nil, errors.New("no flights available")
	}

	maxIndex := -1
	maxSeats := 0

	for i := range flights {
		if flights[i].SeatsAvailable > maxSeats {
			maxSeats = flights[i].SeatsAvailable
			maxIndex = i
		}
	}

	if maxIndex == -1 {
		return nil, errors.New("no available seats on any flight")
	}

	return &flights[maxIndex], nil
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
	}

	// Step 5: Try booking on a non-existent flight
	flights, err = BookSeat(flights, "AA999")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 6: Show updated flight info
	fmt.Println("\nUpdated flight list:")
	ShowAllFlights(flights)

	// Demo Question 10
	firstAvailable, err := GetFirstAvailableFlight(flights)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFirst available flight:", firstAvailable.FlightNumber)
	}

	// Demo Question 11
	fmt.Println("\nTotal available seats:", CountAvailableSeats(flights))

	// Demo Question 12
	mostSeats, err := GetFlightWithMostAvailableSeats(flights)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFlight with most available seats:", mostSeats.FlightNumber)
	}
}
