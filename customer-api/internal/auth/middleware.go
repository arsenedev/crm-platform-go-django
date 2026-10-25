package auth

import (
	"net/http"
	"strings"

	jwtpkg "company/customer-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(
	secret string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader(
			"Authorization",
		)

		if header == "" {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing token",
				},
			)

			return
		}

		tokenString := strings.TrimPrefix(
			header,
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
