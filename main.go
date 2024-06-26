package main

import (
	"context"
	"log"
	"os"

	"github.com/devSobrinho/go-crud/src/configuration/dependencies"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/sse"
	"github.com/devSobrinho/go-crud/src/controller/routes"
	"github.com/devSobrinho/go-crud/src/database/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logger.InitLogger()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	deps := dependencies.InitDependencies(database)

	sse.Stream = sse.NewServer()

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, deps)

	port := ":" + os.Getenv("APP_PORT")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

}
