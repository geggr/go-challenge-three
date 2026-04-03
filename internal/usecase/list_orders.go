package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: orderRepository}
}

func (l *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	output := make([]OrderOutputDTO, 0, len(orders))

	for _, orderItem := range orders {
		output = append(output, OrderOutputDTO{
			ID:         orderItem.ID,
			Price:      orderItem.Price,
			Tax:        orderItem.Tax,
			FinalPrice: orderItem.FinalPrice,
		})
	}
	return output, nil
}
