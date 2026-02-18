package main

import (
	"log"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/Bharat/go-bookstore/middleware"
	"github.com/Bharat/go-bookstore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Logger()
	initializers.Connect()
	initializers.SyncDatabase()
	initializers.InitGoth()
}

func main() {
	defer initializers.Log.Sync()

	r := gin.New()
	r.Use(middleware.ZapLogger())
	r.Use(gin.Recovery())
	routes.RegisterBookstoreRoutes(r)

	log.Println("Starting Gin server on localhost:8080...")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
