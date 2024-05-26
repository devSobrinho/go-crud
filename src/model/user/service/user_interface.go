package service

import (
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"

	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository"
)

func NewUserDomainService(repo repository.UserRepositoryInterface) UserDomainServiceInterface {
	return &userDomainService{
		repo,
	}
}

type userDomainService struct {
	repo repository.UserRepositoryInterface
}

type UserDomainServiceInterface interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdaterUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
