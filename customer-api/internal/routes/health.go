package routes

import (
	"github.com/gin-gonic/gin"
)

func registerHealth(
	r *gin.Engine,
) {

	r.GET("/healthz", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
}
