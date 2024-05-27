package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/auth"
	viewAuth "github.com/devSobrinho/go-crud/src/view/auth"
	viewUser "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (a *authController) LoginUser(c *gin.Context) {
	logger.Info("Inicia LoginUser controller", zap.String("journey", "loginUser"))
	var authRequest request.AuthLoginRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		logger.Error("Erro ao tentar validar as informações do usuário", err, zap.String("journey", "loginUser"))

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info(authRequest.Email, zap.String("journey", "loginUser"))
	logger.Info(authRequest.Password, zap.String("journey", "loginUser"))
	responseUserDomain, err := a.userService.FindUserByEmail(authRequest.Email)
	if err != nil {
		logger.Error("Erro ao tentar chamar LoginUser service", err, zap.String("journey", "loginUser"))

		c.JSON(err.Code, err)
		return
	}

	if err := responseUserDomain.ComparePassword(authRequest.Password); err != nil {
		logger.Error("Erro ao tentar comparar as senhas", err, zap.String("journey", "loginUser"))

		c.JSON(err.Code, err)
		return
	}

	token, err := responseUserDomain.GenerateToken()
	if err != nil {
		logger.Error("Erro ao tentar gerar token", err, zap.String("journey", "loginUser"))

		c.JSON(err.Code, err)
		return
	}

	refreshToken, err := responseUserDomain.GenerateRefreshToken()
	if err != nil {
		logger.Error("Erro ao tentar gerar token", err, zap.String("journey", "loginUser"))

		c.JSON(err.Code, err)
		return
	}

	userResponse := viewUser.ConvertUserDomainToResponse(responseUserDomain)
	response := viewAuth.LoginResponse(token, refreshToken, userResponse)

	c.JSON(http.StatusOK, response)
}
