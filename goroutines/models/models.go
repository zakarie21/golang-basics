package models

import (
	
	"time"
)

type Order struct {
	OrderID string
}

func (o Order) IsShippingAvailable(done chan string, error chan error) {
	time.Sleep(3 * time.Second)
	done<- "Shipping is available for product"
}

func (o Order) IsProductAvailable(done chan string, error chan error) {
	time.Sleep(4 * time.Second)
	done<- "The product is available"
}


func (o Order) IsPaymentSuccessful(done chan string, error chan error) {
	time.Sleep(5 * time.Second)
	
	done<- "Your payment is successful"
}
