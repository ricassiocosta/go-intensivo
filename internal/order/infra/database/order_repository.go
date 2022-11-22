package database

import (
	"database/sql"
	"errors"

	"github.com/ricassiocosta/gointensivo/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) (*OrderRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is missing")
	}

	return &OrderRepository{
		Db: db,
	}, nil
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare(
		"INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}
