package entity

import (
	"errors"
	"strings"
)

var (
	errMissingOrderID         = errors.New("order id is missing")
	errMissingOrderPrice      = errors.New("order price is missing")
	errMissingOrderTax        = errors.New("order tax is missing")
	errMissingOrderFinalPrice = errors.New("order final price is missing")
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}
	if err := order.IsValid(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) IsValid() error {
	o.ID = strings.TrimSpace(o.ID)
	if len(o.ID) == 0 {
		return errMissingOrderID
	}

	if o.Price < 0 {
		return errMissingOrderPrice
	}

	if o.Tax < 0 {
		return errMissingOrderTax
	}

	if o.FinalPrice < 0 {
		return errMissingOrderFinalPrice
	}

	return nil
}
