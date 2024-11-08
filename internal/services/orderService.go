package services

import (
	"HepsiGonulden/internal/repository"
	"HepsiGonulden/internal/types"
	"HepsiGonulden/kafka"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type OrderService struct {
	repo     *repository.OrderRepository
	producer *kafka.Producer
}

func NewOrderService(repo *repository.OrderRepository, producer *kafka.Producer) *OrderService {
	return &OrderService{
		repo:     repo,
		producer: producer,
	}
}

func (s *OrderService) GetById(ctx context.Context, id string) (*types.Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, orderRequestModel *types.OrderRequestModel) (string, error) {
	now := time.Now().Local()
	order := &types.Order{
		Id:            uuid.New().String(),
		OrderName:     orderRequestModel.OrderName,
		OrderTotal:    orderRequestModel.OrderTotal,
		CreatorUserId: orderRequestModel.CreatorUserId,
		CreatedAt:     now,
		PaymentMethod: "Credit Card",
		OrderStatus:   "Getting Ready",
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

func (s *OrderService) Update(ctx context.Context, id string, orderUpdateModel types.OrderUpdateModel) error {
	order, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}

	order.OrderName = orderUpdateModel.OrderName
	order.OrderTotal = orderUpdateModel.OrderTotal
	order.PaymentMethod = orderUpdateModel.PaymentMethod
	order.OrderStatus = orderUpdateModel.OrderStatus
	order.UpdatedAt = time.Now().UTC()
	return s.repo.OrderUpdate(ctx, id, order)
}

func (s *OrderService) Delete(ctx context.Context, id string) error {
	return s.repo.OrderDelete(ctx, id)
}
