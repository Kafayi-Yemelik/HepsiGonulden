package Customer

import (
	"HepsiGonulden/Customer/types"
	"context"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByID(ctx context.Context, id string) (*types.Customer, error) {
	customer, err := s.repo.FindByID(ctx, id)
	if err != nil {

		return nil, err
	}
	return customer, nil
}
func (s *Service) GetByEmail(ctx context.Context, email string) (*types.Customer, error) {
	customer, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err}
	return customer, nil
}
func (s *Service) Create(ctx context.Context, customerRequestModel types.CustomerRequestModel) (string, error) {
	customID := uuid.New().String()
	now := time.Now().Local()
	customerRequestModel.CreatedAt = now

	customer := &types.Customer{
		FirstName: customerRequestModel.FirstName,
		LastName:  customerRequestModel.LastName,
		Age:       customerRequestModel.Age,
		Email:     customerRequestModel.Email,
		CreatedAt: customerRequestModel.CreatedAt,
		Id:        customID,
		Username:  customerRequestModel.Username,
	}

	_, err := s.repo.Create(ctx, customer)
	if err != nil {
		return "", err
	}

	return customID, nil
}
