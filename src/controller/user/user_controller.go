package controller

import (
	"github.com/devSobrinho/go-crud/src/model/user/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	FindUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	service service.UserDomainServiceInterface
}

func NewUserControllerInterface(service service.UserDomainServiceInterface) UserControllerInterface {
	return &userController{
		service: service,
	}
}
