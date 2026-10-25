package auth

import (
	"company/customer-api/internal/audit"
	"company/customer-api/internal/user"
	"company/customer-api/pkg/logger"
	"company/customer-api/pkg/security"

	"context"

	"go.uber.org/zap"
)

type Service struct {
	userRepo     *user.Repository
	jwtSecret    string
	auditService *audit.Service
}

func NewService(
	userRepo *user.Repository,
	jwtSecret string,
	auditService *audit.Service,
) *Service {

	return &Service{
		userRepo:     userRepo,
		jwtSecret:    jwtSecret,
		auditService: auditService,
	}
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (*user.User, error) {

	u, err := s.userRepo.GetByEmail(
		ctx,
		email,
	)

	if err != nil {

		logger.Log.Warn(
			"login_failed_user_not_found",
			zap.String(
				"email",
				email,
			),
		)

		return nil, err
	}

	err = security.CheckPassword(
		u.Password,
		password,
	)

	if err != nil {

		logger.Log.Warn(
			"login_failed_invalid_password",
			zap.String(
				"email",
				email,
			),
		)

		return nil, err
	}

	logger.Log.Info(
		"user_logged_in",
		zap.String(
			"user_id",
			u.ID,
		),
		zap.String(
			"email",
			u.Email,
		),
		zap.String(
			"role",
			u.Role,
		),
	)
	_ = s.auditService.Log(
		ctx,
		u.ID,
		"LOGIN",
		"auth",
		u.ID,
	)

	return u, nil
}
