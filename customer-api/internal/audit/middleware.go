package audit

import (
	"github.com/gin-gonic/gin"
)

func Middleware(
	service *Service,
	action string,
	resource string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		userID := c.GetString(
			"user_id",
		)

		resourceID := c.Param(
			"id",
		)

		_ = service.Log(
			c.Request.Context(),
			userID,
			action,
			resource,
			resourceID,
		)
	}
}
