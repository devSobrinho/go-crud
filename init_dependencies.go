package main

// import (
// 	"github.com/devSobrinho/go-crud/src/configuration/logger"
// 	controller "github.com/devSobrinho/go-crud/src/controller/user"
// 	"github.com/devSobrinho/go-crud/src/model/user/repository"
// 	"github.com/devSobrinho/go-crud/src/model/user/service"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.uber.org/zap"
// )

// func initDependencies(database *mongo.Database) controller.UserControllerInterface {
// 	logger.Info("Iniciando dependências", zap.String("journey", "initDependencies"))
// 	repo := repository.NewUserRepository(database)
// 	userService := service.NewUserDomainService(repo)
// 	userController := controller.NewUserControllerInterface(userService)

// 	return userController
// }
