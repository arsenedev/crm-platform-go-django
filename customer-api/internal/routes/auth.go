package routes

import (
	"company/customer-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func registerAuth(
	r *gin.Engine,
	handler *auth.Handler,
) {

	authGroup := r.Group("/auth")

	authGroup.POST(
		"/login",
		handler.Login,
	)

	/*
		authGroup.POST(
			"/refresh",
			handler.Refresh,
		)

		authGroup.POST(
			"/logout",
			handler.Logout,
		)
	*/
}
