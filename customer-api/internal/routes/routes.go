package routes

import (
	"arsenedev/customer-api/internal/audit"
	"arsenedev/customer-api/internal/auth"
	"arsenedev/customer-api/internal/contact"
	"arsenedev/customer-api/internal/customer"
	"arsenedev/customer-api/internal/user"

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
