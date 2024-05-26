package view

import (
	response "github.com/devSobrinho/go-crud/src/controller/model/response/user"
	model "github.com/devSobrinho/go-crud/src/model/user"
)

func ConvertUserDomainToResponse(userDomain model.UserDomainInterface) response.UserCreateResponse {
	return response.UserCreateResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
