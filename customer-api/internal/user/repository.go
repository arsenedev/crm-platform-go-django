package user

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
	user *User,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO users(
			id,
			email,
			password,
			firstname,
			lastname,
			role,
			active
		)
		VALUES($1,$2,$3,$4,$5,$6,$7)
		`,
		user.ID,
		user.Email,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Role,
		user.Active,
	)

	if err != nil {

		logger.Log.Error(
			"user_create_failed",
			zap.Error(err),
			zap.String("email", user.Email),
		)

		return err
	}

	logger.Log.Info(
		"user_created",
		zap.String("user_id", user.ID),
		zap.String("email", user.Email),
	)

	return nil
}

func (r *Repository) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			password,
			firstname,
			lastname,
			role,
			active,
			created_at
		FROM users
		WHERE email=$1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
	)

	if err != nil {

		logger.Log.Warn(
			"user_not_found",
			zap.String("email", email),
			zap.Error(err),
		)

		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetByID(
	ctx context.Context,
	id string,
) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			firstname,
			lastname,
			role,
			active,
			created_at
		FROM users
		WHERE id=$1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
	)

	if err != nil {

		logger.Log.Error(
			"user_load_failed",
			zap.String("user_id", id),
			zap.Error(err),
		)

		return nil, err
	}

	return &user, nil
}

func (r *Repository) List(
	ctx context.Context,
) ([]User, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			email,
			firstname,
			lastname,
			role,
			active,
			created_at
		FROM users
		ORDER BY created_at DESC
		`,
	)

	if err != nil {

		logger.Log.Error(
			"user_list_failed",
			zap.Error(err),
		)

		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {

		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Firstname,
			&user.Lastname,
			&user.Role,
			&user.Active,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) Update(
	ctx context.Context,
	user *User,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		UPDATE users
		SET
			firstname=$2,
			lastname=$3,
			role=$4,
			active=$5
		WHERE id=$1
		`,
		user.ID,
		user.Firstname,
		user.Lastname,
		user.Role,
		user.Active,
	)

	if err != nil {

		logger.Log.Error(
			"user_update_failed",
			zap.String("user_id", user.ID),
			zap.Error(err),
		)

		return err
	}

	return nil
}

func (r *Repository) Delete(
	ctx context.Context,
	id string,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		DELETE FROM users
		WHERE id=$1
		`,
		id,
	)

	if err != nil {

		logger.Log.Error(
			"user_delete_failed",
			zap.String("user_id", id),
			zap.Error(err),
		)

		return err
	}

	return nil
}
