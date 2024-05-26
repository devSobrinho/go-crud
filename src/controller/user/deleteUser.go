package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *userController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	response, err := u.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Erro ao tentar chamar DeleteUser service", err, zap.String("journey", "createUser"))

		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
