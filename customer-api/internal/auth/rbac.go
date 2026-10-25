package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(
	roles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		value, exists := c.Get(
			"role",
		)

		if !exists {

			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"error": "role missing",
				},
			)

			return
		}

		userRole := value.(string)

		for _, role := range roles {

			if role == userRole {

				c.Next()

				return
			}
		}

		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"error": "access denied",
			},
		)
	}
}
