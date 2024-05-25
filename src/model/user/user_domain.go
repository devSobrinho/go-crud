package model

import (
	"encoding/json"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"go.uber.org/zap"
)

type userDomain struct {
	Id       string `json:"_id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func (u *userDomain) GetID() string {
	return u.Id
}

func (u *userDomain) SetID(id string) {
	u.Id = id
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

func (u *userDomain) GetJSONValue() (string, error) {
	logger.Info("GetJSONValue user", zap.String("journey", "GetJSONValue"))

	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
