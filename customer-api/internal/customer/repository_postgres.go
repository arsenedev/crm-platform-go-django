package customer

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(
	db *pgxpool.Pool,
) Repository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	c *Customer,
) error {

	_, err := r.db.Exec(
		ctx,
		`
        INSERT INTO customers(
            id,
            name,
            email,
            phone,
            status,
            created_at,
            updated_at
        )
        VALUES ($1,$2,$3,$4,$5,NOW(),NOW())
        `,
		c.ID,
		c.Name,
		c.Email,
		c.Phone,
		c.Status,
	)

	return err
}

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*Customer, error) {

	var customer Customer

	err := r.db.QueryRow(
		ctx,
		`
        SELECT
            id,
            name,
            email,
            phone,
            status,
            created_at,
            updated_at
        FROM customers
        WHERE id=$1
        `,
		id,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Status,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]Customer, error) {

	rows, err := r.db.Query(
		ctx,
		`
        SELECT
            id,
            name,
            email,
            phone,
            status,
            created_at,
            updated_at
        FROM customers
        ORDER BY created_at DESC
        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []Customer

	for rows.Next() {

		var c Customer

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.Phone,
			&c.Status,
			&c.CreatedAt,
			&c.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (r *PostgresRepository) Update(
	ctx context.Context,
	c *Customer,
) error {

	_, err := r.db.Exec(
		ctx,
		`
        UPDATE customers
        SET
            name=$2,
            email=$3,
            phone=$4,
            status=$5,
            updated_at=NOW()
        WHERE id=$1
        `,
		c.ID,
		c.Name,
		c.Email,
		c.Phone,
		c.Status,
	)
	// make log entry for update operation

	// Return any error that occurred during the update operation

	return err
}

func (r *PostgresRepository) Delete(
	ctx context.Context,
	id string,
) error {

	_, err := r.db.Exec(
		ctx,
		`DELETE FROM customers WHERE id=$1`,
		id,
	)

	return err
}
