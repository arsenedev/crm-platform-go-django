package user

import (
	"context"

	"company/customer-api/pkg/security"

	"company/customer-api/internal/audit"
	"company/customer-api/pkg/logger"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type Service struct {
	repo         *Repository
	auditService *audit.Service
}

func NewService(
	repo *Repository,
	auditService *audit.Service,
) *Service {

	return &Service{
		repo:         repo,
		auditService: auditService,
	}
}

func (s *Service) Create(
	ctx context.Context,
	req CreateUserRequest,
) error {

	hash, err := security.HashPassword(
		req.Password,
	)

	if err != nil {
		return err
	}

	user := User{
		ID:        uuid.NewString(),
		Email:     req.Email,
		Password:  hash,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Role:      req.Role,
		Active:    true,
	}
	err = s.repo.Create(
		ctx,
		&user,
	)

	if err != nil {

		logger.Log.Error(
			"user_create_failed",
			zap.Error(err),
		)

		return err
	}
	_ = s.auditService.Log(
		ctx,
		user.ID,
		"CREATE_USER",
		"user",
		user.ID,
	)

	logger.Log.Info(
		"user_created",
		zap.String("user_id", user.ID),
		zap.String("email", user.Email),
		zap.String("role", user.Role),
	)

	return nil
}

func (s *Service) List(
	ctx context.Context,
) ([]User, error) {

	users, err := s.repo.List(ctx)

	if err != nil {

		logger.Log.Error(
			"user_list_failed",
			zap.Error(err),
		)

		return nil, err
	}

	logger.Log.Info(
		"users_listed",
		zap.Int("count", len(users)),
	)

	return users, nil

}

func (s *Service) GetByID(
	ctx context.Context,
	id string,
) (*User, error) {
	logger.Log.Info(
		"user_loaded",
		zap.String("user_id", id),
	)

	return s.repo.GetByID(
		ctx,
		id,
	)
}

func (s *Service) Update(
	ctx context.Context,
	id string,
	req UpdateUserRequest,
) error {

	user, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	if req.Firstname != "" {
		user.Firstname = req.Firstname
	}

	if req.Lastname != "" {
		user.Lastname = req.Lastname
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	if req.Active != nil {
		user.Active = *req.Active
	}
	logger.Log.Info(
		"user_updated",
		zap.String("user_id", id),
	)
	err = s.repo.Update(
		ctx,
		user,
	)

	_ = s.auditService.Log(
		ctx,
		user.ID,
		"UPDATE_USER",
		"user",
		user.ID,
	)

	return nil
}

func (s *Service) Delete(
	ctx context.Context,
	id string,
) error {
	logger.Log.Info(
		"user_deleted",
		zap.String(
			"user_id",
			id,
		),
	)
	logger.Log.Info(
		"user_deleted",
		zap.String("user_id", id),
	)
	_ = s.auditService.Log(
		ctx,
		id,
		"DELETE_USER",
		"user",
		id,
	)
	return s.repo.Delete(
		ctx,
		id,
	)

}
