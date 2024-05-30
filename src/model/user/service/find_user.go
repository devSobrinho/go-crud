package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUser(userDomain model.UserDomainInterfacePagination) ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUser", zap.String("journey", "findUser"))
	return u.userRepository.FindUser(userDomain)
}

func (u *userDomainService) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUserById", zap.String("journey", "findUserById"))
	return u.userRepository.FindUserById(id)
}

func (u *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUserByEmail", zap.String("journey", "findUserByEmail"))
	return u.userRepository.FindUserByEmail(email)
}

func (u *userDomainService) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia serviço findUserByEmailAndPassword", zap.String("journey", "findUserByEmailAndPassword"))
	return u.userRepository.FindUserByEmailAndPassword(email, password)
}
