package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("CreateUser model", zap.String("journey", "createUser"))

	return nil
}
