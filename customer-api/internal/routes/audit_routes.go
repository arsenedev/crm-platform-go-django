package routes

import (
	"company/customer-api/internal/audit"
	"company/customer-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuditRoutes(
	api *gin.RouterGroup,
	handler *audit.Handler,
) {

	logs := api.Group("/audit")

	logs.GET(
		"",
		auth.RequireRole("admin"),
		handler.ListLogs,
	)
}
