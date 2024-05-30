package routes

import (
	"github.com/devSobrinho/go-crud/src/configuration/dependencies"
	"github.com/devSobrinho/go-crud/src/configuration/sse"
	"github.com/devSobrinho/go-crud/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, deps dependencies.Dependencies) {

	authorizedRouter := r.Group("/", middleware.Logging)

	userRouter := authorizedRouter.Group("/user")
	{
		userRouter.GET("/", deps.UserController.FindUser)
		userRouter.GET(":userId", deps.UserController.FindUserById)
		userRouter.GET("/email/:userEmail", deps.UserController.FindUserByEmail)
		userRouter.POST("/", deps.UserController.CreateUser)
		userRouter.PUT(":userId", deps.UserController.UpdateUser)
		userRouter.DELETE(":userId", deps.UserController.DeleteUser)
	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", deps.AuthController.LoginUser)
		authRouter.POST("/refresh-token", deps.AuthController.RefreshToken)
	}

	sseRouter := authorizedRouter.Group("/sse")
	{
		sseRouter.GET("/", middleware.SSEMiddleware(), sse.Stream.ServeHTTP(), deps.SSEController.SSE)
		sseRouter.POST("/", deps.SSEController.MessageSSE)
	}

}
