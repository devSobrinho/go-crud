package userController

import (
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserCreateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	c.JSON(200, userRequest)
}
