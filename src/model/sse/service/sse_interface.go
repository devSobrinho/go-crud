package serviceSSE

import "github.com/devSobrinho/go-crud/src/configuration/rest_err"

func NewSSEService() SSEServiceInterface {
	return &sseService{}
}

type sseService struct{}

type SSEServiceInterface interface {
	SendMessage(string, string) (string, *rest_err.RestErr)
}
