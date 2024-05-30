package sseController

import (
	serviceSSE "github.com/devSobrinho/go-crud/src/model/sse/service"
	"github.com/gin-gonic/gin"
)

func NewSSEControllerInterface(serviceSSE serviceSSE.SSEServiceInterface) SSEControllerInterface {
	return &sseControllerStruct{
		serviceSSE: serviceSSE,
	}
}

type sseControllerStruct struct {
	serviceSSE serviceSSE.SSEServiceInterface
}

type SSEControllerInterface interface {
	SSE(c *gin.Context)
	MessageSSE(c *gin.Context)
}
