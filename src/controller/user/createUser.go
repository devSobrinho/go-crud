package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	model "github.com/devSobrinho/go-crud/src/model/user"
	view "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userController) CreateUser(c *gin.Context) {
	logger.Info("Inicia CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserCreateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar as informações do usuário", err, zap.String("journey", "createUser"))

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

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("Erro ao tentar chamar CreateUser service", err, zap.String("journey", "createUser"))

		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertUserDomainToResponse(domainResult)

	logger.Info(
		"CreateUser controller executou com sucesso",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)
	c.JSON(http.StatusCreated, response)
}
