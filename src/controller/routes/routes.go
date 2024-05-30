package routes

import (
	"github.com/devSobrinho/go-crud/src/configuration/dependencies"
	"github.com/devSobrinho/go-crud/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, deps dependencies.Dependencies) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/", middleware.Logging, deps.UserController.FindUser)
		userRouter.GET(":userId", middleware.Logging, deps.UserController.FindUserById)
		userRouter.GET("/email/:userEmail", middleware.Logging, deps.UserController.FindUserByEmail)
		userRouter.POST("/", middleware.Logging, deps.UserController.CreateUser)
		userRouter.PUT(":userId", middleware.Logging, deps.UserController.UpdateUser)
		userRouter.DELETE(":userId", middleware.Logging, deps.UserController.DeleteUser)
	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", deps.AuthController.LoginUser)
		authRouter.POST("/refresh-token", deps.AuthController.RefreshToken)
	}

}
