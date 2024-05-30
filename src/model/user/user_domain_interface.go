package model

import (
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	requestCommon "github.com/devSobrinho/go-crud/src/controller/model/request/common"
)

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(string)
	EncryptPassword() (string, *rest_err.RestErr)
	ComparePassword(string) *rest_err.RestErr
	GenerateToken() (string, *rest_err.RestErr)
	GenerateRefreshToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

// -------------- Pagination --------------
type UserDomainInterfacePagination interface {
	GetUser() UserDomainInterface
	GetPagination() *requestCommon.PaginationQuery
}

type userDomainPagination struct {
	user       *userDomain
	pagination *requestCommon.PaginationQuery
}

func (u *userDomainPagination) GetUser() UserDomainInterface {
	return u.user
}

func (u *userDomainPagination) GetPagination() *requestCommon.PaginationQuery {
	return u.pagination
}

func NewUserDomainPagination(
	email, id string,
	page, size, order string,
) UserDomainInterfacePagination {
	return &userDomainPagination{
		user: &userDomain{
			id:       id,
			email:    email,
			password: "",
			name:     "",
			age:      0,
		},
		pagination: &requestCommon.PaginationQuery{
			Page:  page,
			Size:  size,
			Order: order,
		},
	}
}
