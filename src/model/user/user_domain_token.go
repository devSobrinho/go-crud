package model

import (
	"os"
	"time"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func (u *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(constants.ENV_JWT_SECRET)
	return u.generationToken(secret)
}

func (u *userDomain) GenerateRefreshToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(constants.ENV_JWT_SECRET_REFRESH)
	return u.generationToken(secret)
}

func (u *userDomain) generationToken(secret string) (string, *rest_err.RestErr) {
	claims := jwt.MapClaims{
		"id":    u.GetID(),
		"email": u.GetEmail(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Error("Erro ao gerar token", err, zap.String("journey", "generateToken"))
		return "", rest_err.NewInternalServerError("Erro ao gerar token")
	}

	return tokenString, nil
}
