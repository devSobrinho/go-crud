package model

import rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(string)
	EncryptPassword() (string, *rest_err.RestErr)
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
