package routes

import (
	userController "github.com/devSobrinho/go-crud/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userId", userController.FindUserById)
	r.GET("/user/:userEmail", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
