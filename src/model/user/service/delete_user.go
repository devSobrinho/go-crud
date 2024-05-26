package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (s *userDomainService) DeleteUser(id string) (string, *rest_err.RestErr) {
	logger.Info("DeleteUser model", zap.String("journey", "deleteUser"))

	if _, err := s.userRepository.FindUserById(id); err != nil {
		errRest := rest_err.NewNotFoundError("Usuário não encontrado")
		return "", errRest
	}

	response, err := s.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Erro ao tentar chamar DeleteUser repository", err, zap.String("journey", "deleteUser"))
		return "", err
	}

	return response, nil
}
