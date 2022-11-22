package entity

type IOrderRepository interface {
	Save(order *Order) error
}
