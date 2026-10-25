package audit

import (
	"context"

	"company/customer-api/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(
	db *pgxpool.Pool,
) *Repository {

	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	logEntry AuditLog,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO audit_logs(
			id,
			user_id,
			action,
			resource,
			resource_id
		)
		VALUES($1,$2,$3,$4,$5)
		`,
		logEntry.ID,
		logEntry.UserID,
		logEntry.Action,
		logEntry.Resource,
		logEntry.ResourceID,
	)

	if err != nil {

		logger.Log.Error(
			"audit_insert_failed",
			zap.Error(err),
		)

		return err
	}

	logger.Log.Info(
		"audit_saved",
		zap.String("action", logEntry.Action),
		zap.String("resource", logEntry.Resource),
	)

	return nil
}

func (r *Repository) List(
	ctx context.Context,
) ([]AuditLog, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			user_id,
			action,
			resource,
			resource_id,
			created_at
		FROM audit_logs
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var logs []AuditLog

	for rows.Next() {

		var l AuditLog

		err := rows.Scan(
			&l.ID,
			&l.UserID,
			&l.Action,
			&l.Resource,
			&l.ResourceID,
			&l.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		logs = append(logs, l)
	}

	return logs, nil
}
