// main.go
package main

import (
	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/server/api"
	"github.com/DOSNetwork/DOSscan-api/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	//TODO : Add configuration
	//config.Load()

	models.Connect()
	//defer models.Close()
	//redis.Connect()
	//task.InitEthClient()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(middleware.Limit())
	r.Use(middleware.CORS())
	api.ApplyRoutes(r)
	// Using port :8080 by default
	r.Run()
}
