package contact

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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
	contact *Contact,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO contacts(
			id,
			customer_id,
			firstname,
			lastname,
			email,
			phone
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		`,
		contact.ID,
		contact.CustomerID,
		contact.Firstname,
		contact.Lastname,
		contact.Email,
		contact.Phone,
	)

	return err
}

// List returns all contacts in the database.
func (r *Repository) List(
	ctx context.Context,
) ([]Contact, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			customer_id,
			firstname,
			lastname,
			email,
			phone,
			created_at
		FROM contacts
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact

		if err := rows.Scan(
			&contact.ID,
			&contact.CustomerID,
			&contact.Firstname,
			&contact.Lastname,
			&contact.Email,
			&contact.Phone,
			&contact.CreatedAt,
		); err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}
