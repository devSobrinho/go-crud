package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Erro ao tentar validar o ID do usuário", err, zap.String("journey", "deleteUser"))

		errRest := rest_err.NewBadRequestError("ID do usuário inválido")
		c.JSON(errRest.Code, errRest)
		return
	}

	response, err := u.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Erro ao tentar chamar DeleteUser service", err, zap.String("journey", "createUser"))

		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
