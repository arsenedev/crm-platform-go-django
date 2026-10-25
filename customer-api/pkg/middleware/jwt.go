package middleware

import (
	"net/http"
	"strings"

	jwtpkg "company/customer-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWT(secret string) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader(
			"Authorization",
		)

		if authHeader == "" {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing authorization header",
				},
			)

			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		claims, err := jwtpkg.ValidateToken(
			tokenString,
			secret,
		)

		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			return
		}

		c.Set(
			"user_id",
			claims.UserID,
		)

		c.Set(
			"email",
			claims.Email,
		)

		c.Set(
			"role",
			claims.Role,
		)

		c.Next()
	}
}
