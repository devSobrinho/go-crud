package model

import (
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword() (string, *rest_err.RestErr)
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) EncryptPassword() (string, *rest_err.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	if err != nil {
		return "", rest_err.NewInternalServerError("Error encrypting password")
	}
	logger.Info("EncryptPassword user"+string(hash), zap.String("journey", "encryptPassword"))
	u.password = string(hash)
	return string(hash), nil
}

func (u *userDomain) ComparePassword(hash string) *rest_err.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.password))
	if err != nil {
		return rest_err.NewBadRequestError("Invalid password")
	}

	return nil
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}
