package routes

import (
	"github.com/devSobrinho/go-crud/src/configuration/dependencies"
	"github.com/devSobrinho/go-crud/src/middleware"
	"github.com/gin-gonic/gin"
)

type Gambis interface{}

func InitRoutes(r *gin.RouterGroup, deps dependencies.Dependencies) {
	userRouter := r.Group("/user")
	{
		userRouter.GET(":userId", middleware.Logging, deps.UserController.FindUserById)
		userRouter.GET("/email/:userEmail", deps.UserController.FindUserByEmail)
		userRouter.POST("/", deps.UserController.CreateUser)
		userRouter.PUT(":userId", deps.UserController.UpdateUser)
		userRouter.DELETE(":userId", deps.UserController.DeleteUser)
	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", deps.AuthController.LoginUser)
	}

}
