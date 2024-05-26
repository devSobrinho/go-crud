package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	view "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userController) FindUserById(c *gin.Context) {
	logger.Info("Inicia FindUserById controller", zap.String("journey", "findUserById"))
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Erro ao tentar validar o ID do usuário", err, zap.String("journey", "findUserById"))
		errRest := rest_err.NewBadRequestError("ID do usuário inválido")
		c.JSON(errRest.Code, errRest)
		return
	}

	domainResult, err := u.service.FindUserById(userId)
	if err != nil {
		logger.Error("Erro ao tentar chamar FindUserById service", err, zap.String("journey", "findUserById"))

		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertUserDomainToResponse(domainResult)
	c.JSON(http.StatusOK, response)
}

func (u *userController) FindUserByEmail(c *gin.Context) {
	logger.Info("Inicia FindUserByEmail controller", zap.String("journey", "findUserByEmail"))
	userEmail := c.Param("userEmail")

	domainResult, err := u.service.FindUserByEmail(userEmail)
	if err != nil {
		logger.Error("Erro ao tentar chamar FindUserById service", err, zap.String("journey", "findUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertUserDomainToResponse(domainResult)
	c.JSON(http.StatusOK, response)
}
