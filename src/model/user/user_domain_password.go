package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (u *userDomain) EncryptPassword() (string, *rest_err.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	if err != nil {
		return "", rest_err.NewInternalServerError("Erro ao tentar criptografar senha")
	}
	logger.Info("EncryptPassword user"+string(hash), zap.String("journey", "encryptPassword"))
	u.password = string(hash)
	return string(hash), nil
}

func (u *userDomain) ComparePassword(hash string) *rest_err.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.password))
	if err != nil {
		return rest_err.NewBadRequestError("Senha inválida")
	}

	return nil
}
