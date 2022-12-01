package entity

type IOrderRepository interface {
	Save(order *Order) error
	GetTotal() (int, error)
}
