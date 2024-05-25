package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUser(id string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUser model", zap.String("journey", "findUser"))

	return nil, nil
}
