package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) EncryptPassword() (string, *rest_err.RestErr) {
	password := u.password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", rest_err.NewInternalServerError("Error encrypting password")
	}
	logger.Info("EncryptPassword user"+string(hash), zap.String("journey", "encryptPassword"))

	return string(hash), nil
}

func (u *userDomain) ComparePassword(hash string) *rest_err.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.password))
	if err != nil {
		return rest_err.NewBadRequestError("Invalid password")
	}

	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdaterUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
