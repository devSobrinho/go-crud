package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (s *userDomainService) UpdaterUser(id string, domain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("UpdaterUser model", zap.String("journey", "updaterUser"))
	return s.userRepository.UpdateUser(id, domain)
}
