package controller

import (
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"github.com/gin-gonic/gin"
)

func NewAuthControllerInterface(
	userService service.UserDomainServiceInterface,
) AuthControllerInterface {
	return &authController{
		userService: userService,
	}
}

type AuthControllerInterface interface {
	LoginUser(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authController struct {
	userService service.UserDomainServiceInterface
}
