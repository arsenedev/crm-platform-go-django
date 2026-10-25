package user

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`

	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`

	Role string `json:"role" binding:"required"`
}

type UpdateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
	Active    *bool  `json:"active"`
}
