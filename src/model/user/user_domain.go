package model

import (
	"encoding/json"

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
	GetJSONValue() (string, error)
	SetID(string)
	EncryptPassword() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Id       string `json:"_id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func (u *userDomain) GetJSONValue() (string, error) {
	logger.Info("GetJSONValue user", zap.String("journey", "GetJSONValue"))

	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) SetID(id string) {
	u.Id = id
}

func (u *userDomain) EncryptPassword() (string, *rest_err.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", rest_err.NewInternalServerError("Error encrypting password")
	}
	logger.Info("EncryptPassword user"+string(hash), zap.String("journey", "encryptPassword"))
	u.Password = string(hash)
	return string(hash), nil
}

func (u *userDomain) ComparePassword(hash string) *rest_err.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	if err != nil {
		return rest_err.NewBadRequestError("Invalid password")
	}

	return nil
}
