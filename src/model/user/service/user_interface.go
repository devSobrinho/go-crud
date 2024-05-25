package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"

	model "github.com/devSobrinho/go-crud/src/model/user"
)

func NewUserDomainService() UserDomainServiceInterface {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainServiceInterface interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdaterUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
