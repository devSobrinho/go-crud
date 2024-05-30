package dependencies

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	authController "github.com/devSobrinho/go-crud/src/controller/auth"
	sseController "github.com/devSobrinho/go-crud/src/controller/sse"
	userController "github.com/devSobrinho/go-crud/src/controller/user"
	serviceSSE "github.com/devSobrinho/go-crud/src/model/sse/service"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Dependencies struct {
	UserController userController.UserControllerInterface
	AuthController authController.AuthControllerInterface
	SSEController  sseController.SSEControllerInterface
}

func InitDependencies(database *mongo.Database) Dependencies {
	logger.Info("Iniciando dependências", zap.String("journey", "initDependencies"))
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	userController := userController.NewUserControllerInterface(userService)
	authController := authController.NewAuthControllerInterface(userService)

	serviceSSE := serviceSSE.NewSSEService()
	sseController := sseController.NewSSEControllerInterface(serviceSSE)
	return Dependencies{
		UserController: userController,
		AuthController: authController,
		SSEController:  sseController,
	}
}
