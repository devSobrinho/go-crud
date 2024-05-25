package main

import (
	"log"
	"os"

	"github.com/devSobrinho/go-crud/src/controller/routes"
	controller "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	routes.InitRoutes(&router.RouterGroup, userController)

	port := ":" + os.Getenv("APP_PORT")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

}
