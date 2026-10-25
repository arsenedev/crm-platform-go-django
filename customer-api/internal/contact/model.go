package contact

import "time"

type Contact struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
}
