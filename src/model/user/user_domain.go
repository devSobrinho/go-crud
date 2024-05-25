package model

import (
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (u *UserDomain) EncryptPassword() (string, *rest_err.RestErr) {
	password := u.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", rest_err.NewInternalServerError("Error encrypting password")
	}

	return string(hash), nil
}

func (u *UserDomain) ComparePassword(hash string) *rest_err.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	if err != nil {
		return rest_err.NewBadRequestError("Invalid password")
	}

	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdaterUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
