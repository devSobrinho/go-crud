package sseController

import (
	"fmt"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	requestSSE "github.com/devSobrinho/go-crud/src/controller/model/request/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sseC *sseControllerStruct) MessageSSE(c *gin.Context) {
	var sendMessageRequest requestSSE.SendMessageRequest

	if err := c.ShouldBindJSON(&sendMessageRequest); err != nil {
		logger.Error("Erro ao tentar enviar mensagem", err, zap.String("journey", "messageSSE"))

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	message := sendMessageRequest.Message
	id := sendMessageRequest.ID
	fmt.Println("Id: ", id)
	response, err := sseC.serviceSSE.SendMessage(message, id)
	if err != nil {
		logger.Error("Erro ao tentar enviar mensagem", err, zap.String("journey", "messageSSE"))

		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, gin.H{"message": response})
}
