package api

import (
	"github.com/DOSNetwork/DOSscan-api/server/api/v1"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
