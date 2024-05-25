package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomain) UpdaterUser(id string) *rest_err.RestErr {
	logger.Info("UpdaterUser model", zap.String("journey", "updaterUser"))

	return nil
}
