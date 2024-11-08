package services

import (
	"HepsiGonulden/internal/repository"
	"HepsiGonulden/internal/types"
	"context"
	"github.com/google/uuid"
	"time"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) GetByID(ctx context.Context, id string) (*types.Customer, error) {
	customer, err := s.repo.FindByID(ctx, id)
	if err != nil {

		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) Create(ctx context.Context, customerRequestModel types.CustomerRequestModel) (string, error) {
	customID := uuid.New().String()
	now := time.Now().UTC()

	customer := &types.Customer{
		FirstName:     customerRequestModel.FirstName,
		LastName:      customerRequestModel.LastName,
		Age:           customerRequestModel.Age,
		Email:         customerRequestModel.Email,
		CreatedAt:     now,
		Id:            customID,
		Username:      customerRequestModel.Username,
		Password:      customerRequestModel.Password,
		CreatorUserId: customerRequestModel.CreatorUserId,
	}

	_, err := s.repo.Create(ctx, customer)
	if err != nil {
		return "", err
	}

	return customID, nil
}

func (s *CustomerService) Update(ctx context.Context, id string, customerUpdateModel types.CustomerUpdateModel) error {

	customer, err := s.GetByID(ctx, id)
	now := time.Now().Local()
	if err != nil {
		return err
	}

	customer.FirstName = customerUpdateModel.FirstName
	customer.LastName = customerUpdateModel.LastName
	customer.UpdatedAt = now
	return s.repo.Update(ctx, id, customer)
}
func (s *CustomerService) Delete(ctx context.Context, id string) error {

	return s.repo.Delete(ctx, id)
}

func (s *CustomerService) GetByEmail(ctx context.Context, email string) (*types.Customer, error) {
	customer, err := s.repo.FindByEmail(ctx, email)
	return customer, err
}
