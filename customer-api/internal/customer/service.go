package customer

import (
	"company/customer-api/pkg/logger"
	"context"

	"company/customer-api/internal/audit"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type Service struct {
	repo         Repository
	auditService *audit.Service
}

func NewService(repo Repository, auditService *audit.Service) *Service {
	return &Service{
		repo:         repo,
		auditService: auditService,
	}
}

func (s *Service) Create(
	ctx context.Context,
	req CreateCustomerRequest,
) error {

	customer := Customer{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Status: "active",
	}
	logger.Log.Info(
		"customer_created",
		zap.String("customer_id", customer.ID),
		zap.String("name", customer.Name),
		zap.String("email", customer.Email),
	)

	err := s.repo.Create(ctx, &customer)
	if err != nil {
		logger.Log.Error(
			"customer_creation_failed",
			zap.Error(err),
			zap.String("customer_id", customer.ID),
		)
		return err
	}
	_ = s.auditService.Log(
		ctx,
		"system",
		"CREATE_CUSTOMER",
		"customer",
		customer.ID,
	)

	return nil
}

func (s *Service) List(
	ctx context.Context,
) ([]Customer, error) {

	customers, err := s.repo.List(ctx)

	if err != nil {

		logger.Log.Error(
			"customer_list_failed",
			zap.Error(err),
		)

		return nil, err
	}

	logger.Log.Info(
		"customers_listed",
		zap.Int("count", len(customers)),
	)

	return customers, nil

}

func (s *Service) GetByID(
	ctx context.Context,
	id string,
) (*Customer, error) {
	logger.Log.Info(
		"customer_loaded",
		zap.String("customer_id", id),
	)
	_ = s.auditService.Log(
		ctx,
		"system",
		"GetByID_CUSTOMER",
		"customer",
		id,
	)
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(
	ctx context.Context,
	id string,
	req UpdateCustomerRequest,
) error {

	customer, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if req.Name != "" {
		customer.Name = req.Name
	}

	if req.Email != "" {
		customer.Email = req.Email
	}

	if req.Phone != "" {
		customer.Phone = req.Phone
	}

	if req.Status != "" {
		customer.Status = req.Status
	}
	logger.Log.Info(
		"customer_updated",
		zap.String("customer_id", id),
	)
	_ = s.auditService.Log(
		ctx,
		"system",
		"UPDATE_CUSTOMER",
		"customer",
		customer.ID,
	)
	return s.repo.Update(ctx, customer)
}

func (s *Service) Delete(
	ctx context.Context,
	id string,
) error {
	logger.Log.Info(
		"customer_deleted",
		zap.String("customer_id", id),
	)
	_ = s.auditService.Log(
		ctx,
		"system",
		"DELETE_CUSTOMER",
		"customer",
		id,
	)
	return s.repo.Delete(ctx, id)
}
