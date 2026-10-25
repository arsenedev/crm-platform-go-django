package routes

import (
	"company/customer-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func registerProtected(
	r *gin.Engine,
	jwtSecret string,
	h Handlers,
) {

	api := r.Group("/api/v1")

	api.Use(
		auth.JWTMiddleware(
			jwtSecret,
		),
	)

	registerCustomers(
		api,
		h,
	)

	registerContacts(
		api,
		h,
	)
	RegisterUserRoutes(
		api,
		h.User,
	)
	RegisterAuditRoutes(
		api,
		h.Audit,
	)
}
