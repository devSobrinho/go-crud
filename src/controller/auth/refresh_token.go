package controller

import (
	"net/http"
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/auth"
	"github.com/devSobrinho/go-crud/src/utils"
	viewAuth "github.com/devSobrinho/go-crud/src/view/auth"
	viewUser "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (a *authController) RefreshToken(c *gin.Context) {
	logger.Info("Inicia RefreshToken controller", zap.String("journey", "refreshToken"))
	var refreshTokenRequest request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		logger.Error("Erro ao tentar validar as informações do refreshToken", err, zap.String("journey", "refreshToken"))

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	secretKeyRefreshToken := os.Getenv(constants.ENV_JWT_SECRET_REFRESH)
	claims, err := utils.ValidationTokenAndExtractClaims(secretKeyRefreshToken, refreshTokenRequest.RefreshToken)
	if err != nil {
		logger.Error("Erro ao tentar validar token", err, zap.String("journey", "refreshToken"))

		c.JSON(err.Code, err)
		return
	}

	id := claims["id"].(string)

	userDomain, err := a.userService.FindUserById(id)
	if err != nil {
		logger.Error("Erro ao tentar buscar usuário", err, zap.String("journey", "refreshToken"))

		c.JSON(err.Code, err)
		return
	}
	token, err := userDomain.GenerateToken()
	if err != nil {
		logger.Error("Erro ao tentar gerar token", err, zap.String("journey", "refreshToken"))

		c.JSON(err.Code, err)
		return
	}

	refreshToken, err := userDomain.GenerateRefreshToken()
	if err != nil {
		logger.Error("Erro ao tentar gerar refreshToken", err, zap.String("journey", "refreshToken"))

		c.JSON(err.Code, err)
		return
	}

	user := viewUser.ConvertUserDomainToResponse(userDomain)
	response := viewAuth.LoginResponse(token, refreshToken, user)

	c.JSON(http.StatusOK, response)
}
