package entity

import (
	"errors"
	"strings"
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
		return errors.New("ID is missing")
	}

	if o.Price < 0 {
		return errors.New("price is invalid")
	}

	if o.Tax < 0 {
		return errors.New("tax is invalid")
	}

	if o.FinalPrice < 0 {
		return errors.New("final price is invalid")
	}

	return nil
}
