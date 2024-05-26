package routes

import (
	"github.com/devSobrinho/go-crud/src/configuration/dependencies"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, deps dependencies.Dependencies) {
	r.GET("/user/:userId", deps.UserController.FindUserById)
	r.GET("/user/email/:userEmail", deps.UserController.FindUserByEmail)
	r.POST("/user", deps.UserController.CreateUser)
	r.PUT("/user/:userId", deps.UserController.UpdateUser)
	r.DELETE("/user/:userId", deps.UserController.DeleteUser)

	r.POST("/auth/login", deps.AuthController.LoginUser)
}
