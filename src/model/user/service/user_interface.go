package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"

	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
)

func NewUserDomainService(userRepository repository.UserRepositoryInterface) UserDomainServiceInterface {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

type UserDomainServiceInterface interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdaterUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
