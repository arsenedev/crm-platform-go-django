package contact

type CreateContactRequest struct {
	CustomerID string `json:"customer_id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}
