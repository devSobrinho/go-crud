package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("DeleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
