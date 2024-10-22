package order

import (
	"HepsiGonulden/customer"

	"HepsiGonulden/order/types"
	"context"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo    *Repository
	service *customer.Service
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) GetById(ctx context.Context, id string) (*types.Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {

		return nil, err
	}
	return order, nil
}

func (s *Service) CreateOrder(ctx context.Context, orderRequestModel *types.OrderRequestModel) (string, error) {

	now := time.Now().Local()
	order := &types.Order{
		Id:         uuid.New().String(),
		CustomerId: orderRequestModel.CustomerId,
		OrderName:  orderRequestModel.OrderName,
		OrderTotal: orderRequestModel.OrderTotal,
		CreatedAt:  now,
	}
	_, err := s.repo.OrderCreate(ctx, order)
	if err != nil {
		return "", err
	}

	return order.Id, nil

}
func (s *Service) Update(ctx context.Context, id string, orderUpdateModel types.OrderUpdateModel) error {

	order, err := s.GetById(ctx, id)
	now := time.Now().Local()
	if err != nil {
		return err
	}

	order.OrderName = orderUpdateModel.OrderName
	order.OrderTotal = orderUpdateModel.OrderTotal
	order.UpdatedAt = now
	return s.repo.OrderUpdate(ctx, id, order)
}
func (s *Service) Delete(ctx context.Context, id string) error {

	return s.repo.OrderDelete(ctx, id)
}
