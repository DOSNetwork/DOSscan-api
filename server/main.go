// main.go
package main

import (
	"github.com/DOSNetwork/DOSscan-api/server/handler"
	"github.com/DOSNetwork/DOSscan-api/server/middleware"
	"github.com/DOSNetwork/DOSscan-api/server/repository"

	"github.com/gin-gonic/gin"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

func main() {
	//TODO : Add configuration
	//config.Load()

	db := repository.Connect(DB_USER, DB_PASSWORD, DB_NAME)
	eventsRepo := repository.NewDBEventsRepository(db)
	searchHandler := handler.NesSearchHandler(eventsRepo)
	searchHandler.Init()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(middleware.CORS())

	//Set api route
	api := r.Group("/api")
	v1 := api.Group("/explorer")
	v1.GET("/search", searchHandler.Search)

	// Using port :8080 by default
	r.Run()
}
