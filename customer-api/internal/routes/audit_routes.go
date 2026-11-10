package routes

import (
	"arsenedev/customer-api/internal/audit"
	"arsenedev/customer-api/internal/auth"

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
