package audit

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {

	return &Service{
		repo: repo,
	}
}

func (s *Service) Log(
	ctx context.Context,
	userID string,
	action string,
	resource string,
	resourceID string,
) error {

	return s.repo.Create(
		ctx,
		AuditLog{
			ID:         uuid.NewString(),
			UserID:     userID,
			Action:     action,
			Resource:   resource,
			ResourceID: resourceID,
		},
	)
}

func (s *Service) List(
	ctx context.Context,
) ([]AuditLog, error) {

	return s.repo.List(ctx)
}
