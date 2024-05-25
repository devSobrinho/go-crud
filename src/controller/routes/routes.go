package routes

import (
	controller "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	r.GET("/user/:userId", userController.FindUserById)
	r.GET("/user/email/:userEmail", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
