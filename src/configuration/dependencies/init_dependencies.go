package dependencies

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	authController "github.com/devSobrinho/go-crud/src/controller/auth"
	userController "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Dependencies struct {
	UserController userController.UserControllerInterface
	AuthController authController.AuthControllerInterface
}

func InitDependencies(database *mongo.Database) Dependencies {
	logger.Info("Iniciando dependências", zap.String("journey", "initDependencies"))
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	userController := userController.NewUserControllerInterface(userService)
	authController := authController.NewAuthControllerInterface(userService)
	return Dependencies{
		UserController: userController,
		AuthController: authController,
	}
}
