package usecase

import "github.com/ricassiocosta/gointensivo/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.IOrderRepository
}

func NewGetTotalUseCase(orderRepository entity.IOrderRepository) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
