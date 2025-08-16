package models

import "fmt"

type CreditCard struct{
	CardNumber string
	
}

func (c CreditCard) Pay(amount float64) string{
	return fmt.Sprintf("You payed %0.2f using %s", amount, c.CardNumber)
}