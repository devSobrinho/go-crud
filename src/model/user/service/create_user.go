package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (service *userDomainService) CreateUser(domain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("CreateUser model", zap.String("journey", "createUser"))
	if _, err := domain.EncryptPassword(); err != nil {
		return err
	}
	return nil
}
