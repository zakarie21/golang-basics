package main

import "fmt"

// --------------------
// Patient struct
// --------------------
type Patient struct {
	Name        string
	Age         int
	HealthScore int
}

// Constructor function for Patient
func NewPatient(name string, age int) Patient {
	return Patient{
		Name:        name,
		Age:         age,
		HealthScore: 100, // default full health
	}
}

// Method to update health (pointer receiver needed)
func (p *Patient) UpdateHealth(score int) {
	p.HealthScore = score
}

// Implement Greetable interface
func (p Patient) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", p.Name)
}

// --------------------
// Doctor struct
// --------------------
type Doctor struct {
	Name      string
	Specialty string
}

// Implement Greetable interface
func (d Doctor) Greet() string {
	return fmt.Sprintf("Hi, I am Dr. %s, specialty: %s", d.Name, d.Specialty)
}

// --------------------
// Hospital struct
// --------------------
type Hospital struct {
	Patients []Patient
}

// Constructor for Hospital
func NewHospital() Hospital {
	return Hospital{
		Patients: []Patient{}, // start empty
	}
}

// Method to add a patient
func (h *Hospital) AddPatient(p Patient) {
	h.Patients = append(h.Patients, p)
}

// --------------------
// Interface
// --------------------
type Greetable interface {
	Greet() string
}

// Function to print greetings
func Introduce(g Greetable) {
	fmt.Println(g.Greet())
}

// --------------------
// Main
// --------------------
func main() {
	// Create hospital
	hospital := NewHospital()

	// Create patients
	p1 := NewPatient("Alice", 30)
	p2 := NewPatient("Bob", 45)

	// Add patients to hospital
	hospital.AddPatient(p1)
	hospital.AddPatient(p2)

	// Update health score of Alice
	hospital.Patients[0].UpdateHealth(80)

	// Print patient info
	for _, patient := range hospital.Patients {
		fmt.Printf("Patient: %s, Age: %d, Health: %d\n",
			patient.Name, patient.Age, patient.HealthScore)
	}

	// Create a doctor
	doc := Doctor{Name: "Smith", Specialty: "Cardiology"}

	// Introduce patients and doctor
	fmt.Println("\nGreetings:")
	Introduce(p1)
	Introduce(p2)
	Introduce(doc)
}