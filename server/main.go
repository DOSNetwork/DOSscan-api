// main.go
package main

import (
	"github.com/DOSNetwork/DOSscan-api/server/handler"
	"github.com/DOSNetwork/DOSscan-api/server/middleware"
	"github.com/DOSNetwork/DOSscan-api/server/repository"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/contrib/static"

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
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	eventsRepo := repository.NewDBEventsRepository(db, client)
	searchHandler := handler.NesSearchHandler(eventsRepo)
	searchHandler.Init()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(middleware.CORS())

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./view", true)))
	r.Use(static.Serve("/explorer", static.LocalFile("./view", true)))
	r.Use(static.Serve("/myaccount", static.LocalFile("./view", true)))
	r.Use(static.Serve("/nodelist", static.LocalFile("./view", true)))

	//Set api route
	api := r.Group("/api")
	v1 := api.Group("/explorer")
	v1.GET("/search", searchHandler.Search)

	// Using port :8080 by default
	r.Run()
}
