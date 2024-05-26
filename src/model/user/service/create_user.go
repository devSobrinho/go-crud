package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (service *userDomainService) CreateUser(domain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia CreateUser service", zap.String("journey", "createUser"))
	if _, err := domain.EncryptPassword(); err != nil {
		logger.Error("Erro ao tentar chamar encriptar a senha", err, zap.String("journey", "createUser"))
		return nil, err
	}

	repository, err := service.userRepository.CreateUser(domain)

	if err != nil {
		logger.Error("Erro ao tentar chamar o repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executado com sucesso", zap.String("userId", repository.GetID()), zap.String("journey", "createUser"))

	return repository, nil
}
