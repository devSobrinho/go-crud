package main

import (
	"log"
	"os"

	"github.com/devSobrinho/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	port := ":" + os.Getenv("APP_PORT")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

}
