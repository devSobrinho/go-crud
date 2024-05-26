package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userController) UpdateUser(c *gin.Context) {
	logger.Info("Inicia UpdateUser controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar as informações do usuário", err, zap.String("journey", "updateUser"))

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Erro ao tentar validar o ID do usuário", err, zap.String("journey", "updateUser"))

		errRest := rest_err.NewBadRequestError("ID do usuário inválido")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := u.service.UpdaterUser(userId, domain)
	if err != nil {
		logger.Error("Erro ao tentar chamar UpdateUser service", err, zap.String("journey", "updateUser"))

		c.JSON(err.Code, err)
		return
	}

	responseMessage := "Usuário atualizado com sucesso"
	c.JSON(http.StatusOK, responseMessage)
}
