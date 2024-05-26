package viewAuth

import (
	response "github.com/devSobrinho/go-crud/src/controller/model/response/user"
)

type LoginResponseDto struct {
	Token string                      `json:"token"`
	User  response.UserCreateResponse `json:"user"`
}

func LoginResponse(token string, user response.UserCreateResponse) LoginResponseDto {
	return LoginResponseDto{
		Token: token,
		User:  user,
	}
}
