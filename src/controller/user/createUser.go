package controller

import (
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	model "github.com/devSobrinho/go-crud/src/model/user"
	view "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userController) CreateUser(c *gin.Context) {

	var userRequest request.UserCreateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	response := view.ConvertUserDomainToResponse(domain)

	c.JSON(201, response)
}
