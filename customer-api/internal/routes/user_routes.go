package routes

import (
	"company/customer-api/internal/auth"
	"company/customer-api/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(
	api *gin.RouterGroup,
	handler *user.Handler,
) {

	users := api.Group("/users")

	users.GET(
		"",
		auth.RequireRole("admin"),
		handler.ListUsers,
	)

	users.GET(
		"/:id",
		auth.RequireRole("admin"),
		handler.GetUser,
	)

	users.POST(
		"",
		auth.RequireRole("admin"),
		handler.CreateUser,
	)

	users.PUT(
		"/:id",
		auth.RequireRole("admin"),
		handler.UpdateUser,
	)

	users.DELETE(
		"/:id",
		auth.RequireRole("admin"),
		handler.DeleteUser,
	)
}
