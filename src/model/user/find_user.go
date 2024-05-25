package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *UserDomain) FindUser(id string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("FindUser model", zap.String("journey", "findUser"))

	return u, nil
}
