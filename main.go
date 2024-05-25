package main

import (
	"context"
	"log"
	"os"

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

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	port := ":" + os.Getenv("APP_PORT")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

}
