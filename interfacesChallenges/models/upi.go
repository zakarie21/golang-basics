package models

import "fmt"

type Upi struct{
	UpiNumber string
}

func (u Upi) Pay(amount float64) string{
	return fmt.Sprintf("You paid %.2f using this UPI %s", amount, u.UpiNumber)
}