package order

import (
	"HepsiGonulden/kafka"
	"HepsiGonulden/order/types"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo     *Repository
	producer *kafka.Producer
}

func NewService(repo *Repository, producer *kafka.Producer) *Service {
	return &Service{
		repo:     repo,
		producer: producer,
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
		Id:            uuid.New().String(),
		OrderName:     orderRequestModel.OrderName,
		OrderTotal:    orderRequestModel.OrderTotal,
		CreatorUserId: orderRequestModel.CreatorUserId,
		CreatedAt:     now,
	}
	_, err := s.repo.OrderCreate(ctx, order)
	if err != nil {
		return "", err
	}

	err = s.producer.Publish("order_create", order)
	if err != nil {
		fmt.Printf("kafka order create message produce failed, err: %s", err.Error())
	}

	return order.Id, nil
}

func (s *Service) Update(ctx context.Context, id string, orderUpdateModel types.OrderUpdateModel) error {
	order, err := s.GetById(ctx, id)
	now := time.Now().UTC()
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
