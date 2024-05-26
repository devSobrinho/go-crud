package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUser(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUser", zap.String("journey", "findUser"))

	return nil, nil
}

func (u *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUserByEmail", zap.String("journey", "findUserByEmail"))
	return u.userRepository.FindUserByEmail(email)
}
