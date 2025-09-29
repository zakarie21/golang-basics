package main

import (
	"errors"
	"fmt"
)

type Patient struct {
	Name        string
	Age         int
	HealthScore int
}

func NewPatient(name string, age int) Patient {
	return Patient{
		Name:        name,
		Age:         age,
		HealthScore: 100,
	}
}

func (p *Patient) Treat(amount int) error {
	newScore := p.HealthScore - amount
	if newScore < 0 {
		return errors.New("patient health too low")
	}
	p.HealthScore = newScore
	return nil
}

type Doctor struct {
	Name      string
	Specialty string
}

type Hospital struct {
	Patients []Patient
	Doctor   *Doctor
}

func NewHospital() Hospital {
	return Hospital{
		Patients: []Patient{},
		Doctor:   nil,
	}
}

func (h *Hospital) AddPatient(p Patient) error {
	for _, existing := range h.Patients {
		if existing.Name == p.Name {
			return errors.New("patient already exists")
		}
	}
	h.Patients = append(h.Patients, p)
	return nil
}

func (h *Hospital) AssignDoctor(d Doctor) error {
	if h.Doctor != nil {
		return errors.New("doctor already assigned")
	}
	h.Doctor = &d
	return nil
}

func (h *Hospital) FindPatient(name string) (*Patient, error) {
	for i := range h.Patients {
		if h.Patients[i].Name == name {
			return &h.Patients[i], nil
		}
	}
	return nil, errors.New("patient not found")
}

func (h *Hospital) DischargePatient(name string) error {
	for i, p := range h.Patients {
		if p.Name == name {
			h.Patients = append(h.Patients[:i], h.Patients[i+1:]...)
			return nil
		}
	}
	return errors.New("patient not in hospital")
}

func main() {
	h := NewHospital()

	p1 := NewPatient("Alice", 30)
	p2 := NewPatient("Bob", 40)

	err := h.AddPatient(p1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = h.AddPatient(p2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = h.AddPatient(p1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = h.Patients[0].Treat(50)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = h.Patients[0].Treat(60)
	if err != nil {
		fmt.Println("Error:", err)
	}

	d := Doctor{Name: "Dr. Smith", Specialty: "Cardiology"}
	err = h.AssignDoctor(d)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = h.AssignDoctor(d)
	if err != nil {
		fmt.Println("Error:", err)
	}

	p, err := h.FindPatient("Alice")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found patient:", p.Name, "Health:", p.HealthScore)
	}

	err = h.DischargePatient("Alice")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = h.DischargePatient("Alice")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
