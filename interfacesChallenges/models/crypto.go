package models

import "fmt"

type CryptoWallet struct{
	ID string
}

func (c CryptoWallet) Pay(amount float64) string{
	return fmt.Sprintf("You paid %.2f using crypto wallet %s", amount, c.ID)
}