package main

import (
	"context"
	"log"
	"os"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/controller/routes"
	controller "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/devSobrinho/go-crud/src/database/mongodb"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	logger.Info("Iniciando dependências", zap.String("journey", "initDependencies"))
	repo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(userService)

	return userController
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	// userRepository := repository.NewUserRepository(database)
	// userService := service.NewUserDomainService(userRepository)
	// userController := controller.NewUserControllerInterface(userService)

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	port := ":" + os.Getenv("APP_PORT")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

}
