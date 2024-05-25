package main

import (
	controller "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(userService)

	return userController
}
