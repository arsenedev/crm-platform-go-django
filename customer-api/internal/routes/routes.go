package routes

import (
	"company/customer-api/internal/audit"
	"company/customer-api/internal/auth"
	"company/customer-api/internal/contact"
	"company/customer-api/internal/customer"
	"company/customer-api/internal/user"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth     *auth.Handler
	Customer *customer.Handler
	Contact  *contact.Handler
	User     *user.Handler
	Audit    *audit.Handler
}

func Register(
	r *gin.Engine,
	jwtSecret string,
	h Handlers,
) {

	registerHealth(r)

	registerAuth(
		r,
		h.Auth,
	)

	registerProtected(
		r,
		jwtSecret,
		h,
	)

}
