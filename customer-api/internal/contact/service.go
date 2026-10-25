package contact

import (
	"company/customer-api/pkg/logger"
	"context"

	"company/customer-api/internal/audit"

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
	req CreateContactRequest,
) error {

	contact := Contact{
		ID:         uuid.NewString(),
		CustomerID: req.CustomerID,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Email:      req.Email,
		Phone:      req.Phone,
	}
	logger.Log.Info(
		"contact_created",
		zap.String("contact_id", contact.ID),
		zap.String("customer_id", contact.CustomerID),
		zap.String("email", contact.Email),
	)
	_ = s.auditService.Log(
		ctx,
		"system",
		"CREATE_CONTACT",
		"contact",
		contact.ID,
	)

	return s.repo.Create(
		ctx,
		&contact,
	)
}

func (s *Service) List(
	ctx context.Context,
) ([]Contact, error) {
	contacts, err := s.repo.List(ctx)

	if err != nil {

		logger.Log.Error(
			"contact_list_failed",
			zap.Error(err),
		)

		return nil, err
	}

	logger.Log.Info(
		"contacts_listed",
		zap.Int("count", len(contacts)),
	)

	return contacts, nil
}
